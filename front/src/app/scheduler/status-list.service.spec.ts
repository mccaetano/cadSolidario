import { TestBed } from '@angular/core/testing';

import { StatusListService } from './status-list.service';

describe('StatusListService', () => {
  let service: StatusListService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(StatusListService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
