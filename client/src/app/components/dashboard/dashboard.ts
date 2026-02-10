import { ChangeDetectorRef, Component, inject, OnInit } from '@angular/core';
import { FormArray, FormBuilder, FormGroup, ReactiveFormsModule } from '@angular/forms';
import { AuthService } from '../../services/auth';
import { Claims } from '../../models/auth.model';
import { CheckIns, NewCheckIns } from '../../models/check-in.model';
import { CheckInService } from '../../services/check-in';
import { Error } from '../../models/error.model';
import { Member } from '../../models/team.model';
import { TeamService } from '../../services/team';

// Visibility value stored in the form
type VisibilityValue = 'all' | 'private' | number;

@Component({
  selector: 'app-dashboard',
  imports: [ReactiveFormsModule],
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
    this.auth.userID$.subscribe(uid => {
      this.uid = Number(uid);
      this.loadToday();
    })
  }

  loadToday(): void {
    this.ciSvc.getDaily(this.uid).subscribe({
      next: (resp: CheckIns | null) => {
        this.ci = resp;

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

    let checkIns: NewCheckIns[] = [];

    const sections: { type: 'yesterday' | 'today' | 'blockers'; array: FormArray }[] = [
      { type: 'yesterday', array: this.yesterday },
      { type: 'today', array: this.today },
      { type: 'blockers', array: this.blockers },
    ];

    sections.forEach(section => {
      section.array.controls.forEach(control => {
        const value = control.value;
        // Only include if item or jira is non-empty
        if ((value.item?.trim() ?? '') !== '' || (value.jira?.trim() ?? '') !== '') {
          let vis: string | number;
          vis = value.visibility === "all" || value.visibility === "private" ? value.visibility : parseInt(value.visibility);
          checkIns.push({
            userID: this.uid,
            type: section.type,
            item: value.item?.trim() ?? '',
            jira: value.jira?.trim() || undefined,
            visibility: vis
          });
        }
      });
    });

    if (checkIns.length === 0) {
      alert('Please fill at least one entry.');
      return;
    }

    let payload = { checkIns };

    console.log('Payload:', payload);

    // TODO: wire to API
    // this.ciSvc.submit(payload).subscribe(...)
  }

}