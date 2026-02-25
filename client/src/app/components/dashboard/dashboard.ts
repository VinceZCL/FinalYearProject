import { ChangeDetectorRef, Component, inject, OnInit } from '@angular/core';
import { FormArray, FormBuilder, FormGroup, ReactiveFormsModule } from '@angular/forms';
import { AuthService } from '../../services/auth';
import { CheckIns, NewCheckIns } from '../../models/check-in.model';
import { CheckInService } from '../../services/check-in';
import { Error } from '../../models/error.model';
import { Member } from '../../models/team.model';
import { TeamService } from '../../services/team';
import { DatePipe } from '@angular/common';

// Visibility value stored in the form
type VisibilityValue = 'all' | 'private' | number;

@Component({
  selector: 'app-dashboard',
  imports: [ReactiveFormsModule, DatePipe],
  templateUrl: './dashboard.html',
  styleUrl: './dashboard.css',
})
export class Dashboard implements OnInit {

  form!: FormGroup;
  uid!: number;
  ci!: CheckIns | null;

  private fb = inject(FormBuilder);
  private auth = inject(AuthService);
  private ciSvc = inject(CheckInService);
  private teamSvc = inject(TeamService);
  private cd = inject(ChangeDetectorRef);

  visibilityOptions: { label: string; value: VisibilityValue }[] = [
    { label: 'All', value: 'all' },
    { label: 'Private', value: 'private' },
  ];

  ngOnInit(): void {

    this.auth.claim$.subscribe(claims => {
      this.uid = Number(claims?.userID);
      this.loadToday();
    });

  }

  loadToday(): void {
    this.ciSvc.getDaily(this.uid).subscribe({
      next: (resp: CheckIns | null) => {
        this.ci = resp;
        this.cd.detectChanges();

        if (!this.ci) {
          this.teamSvc.getOwnTeams(this.uid).subscribe({
            next: (val: Member[]) => {
              let teamOpts = val.map(team => ({
                label: team.team_name,
                value: team.teamID,
              }));
              this.visibilityOptions = [
                ...this.visibilityOptions,
                ...teamOpts
              ];
              this.initForm();
              this.cd.detectChanges();
            }
          });
        }
      },
      error: (err: Error) => {
        console.log(err.error);
      }
    });
  }

  initForm(): void {
    this.form = this.fb.group({
      yesterday: this.fb.array([]),
      today: this.fb.array([]),
      blockers: this.fb.array([])
    });

    this.initSection(this.yesterday);
    this.initSection(this.today);
    this.initSection(this.blockers);
  }

  get yesterday(): FormArray {
    return this.form.get('yesterday') as FormArray;
  }

  get today(): FormArray {
    return this.form.get('today') as FormArray;
  }

  get blockers(): FormArray {
    return this.form.get('blockers') as FormArray;
  }

  initSection(section: FormArray): void {
    section.push(this.createPair(section));
  }

  createPair(section: FormArray): FormGroup {
    const group = this.fb.group({
      item: [''],
      jira: [''],
      visibility: ['all' as VisibilityValue]
    });

    let triggerUpdate = () => this.updateInputs(section);

    group.get('item')?.valueChanges.subscribe(triggerUpdate);
    group.get('jira')?.valueChanges.subscribe(triggerUpdate);

    return group;
  }

  updateInputs(section: FormArray): void {
    const controls = section.controls as FormGroup[];
    const last = controls[controls.length - 1];

    const lastItem = last.get('item')?.value?.trim();
    const lastJira = last.get('jira')?.value?.trim();
    const hasLast = lastItem !== '' || lastJira !== '';

    if (hasLast) {
      section.push(this.createPair(section));
    }

    // Remove extra empty rows
    while (
      section.length > 1 &&
      this.isEmpty(controls[controls.length - 1]) &&
      this.isEmpty(controls[controls.length - 2])
    ) {
      section.removeAt(section.length - 1);
    }
  }

  // helper
  isEmpty(group: FormGroup): boolean {
    return (
      group.get('item')?.value?.trim() === '' &&
      group.get('jira')?.value?.trim() === ''
    );
  }

  submit(): void {
    if (!this.form) return;

    const sections: { type: 'yesterday' | 'today' | 'blockers'; array: FormArray }[] = [
      { type: 'yesterday', array: this.yesterday },
      { type: 'today', array: this.today },
      { type: 'blockers', array: this.blockers },
    ];

    const checkIns: NewCheckIns[] = sections.flatMap(section =>
      section.array.controls
        .map(control => {
          const value = control.value;
          const item = value.item?.trim() ?? '';
          const jira = value.jira?.trim() || undefined;

          // Skip if both item and jira are empty
          if (!item && !jira) return null;

          // Determine visibility
          let vis: string | number = value.visibility;
          const parsedTeamID = parseInt(value.visibility);
          if (!isNaN(parsedTeamID)) {
            vis = 'team';
          }

          const checkIn: NewCheckIns = {
            userID: this.uid,
            type: section.type,
            item,
            jira,
            visibility: vis as 'all' | 'private' | 'team',
            teamID: vis === 'team' ? parsedTeamID : undefined
          };

          return checkIn;
        })
        .filter(Boolean) as NewCheckIns[]
    );

    if (checkIns.length === 0) {
      alert('Please fill at least one entry.');
      return;
    }

    const payload = { checkIns };
    this.ciSvc.submitBulk(payload).subscribe({
      next: (resp: CheckIns) => {
        this.ci = resp;
        this.cd.detectChanges();
      },
      error: (err: Error) => {
        console.error(err);
      }
    });
  }

  // TODO also handle get yesterday for suggestions
  // TODO consider get previous (latest checkin)
  // TODO ^ new backend API / reuse date checkin

  // TODO check if checkin exist for today, then display it
  // TODO show date selector to change date, query for checkin of that date
  // ! should you be able to see previous, if not checkin today
  // TODO if no checkin on that day, say not found, show the date of that day as well

}