import { TestBed, inject } from '@angular/core/testing';

import { PermissionCheckService } from './permission-check.service';

describe('PermissionCheckService', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [PermissionCheckService]
    });
  });

  it('should be created', inject([PermissionCheckService], (service: PermissionCheckService) => {
    expect(service).toBeTruthy();
  }));
});
