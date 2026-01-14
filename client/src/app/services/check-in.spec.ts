import { TestBed } from '@angular/core/testing';

import { CheckIn } from './check-in';

describe('CheckIn', () => {
  let service: CheckIn;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(CheckIn);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
