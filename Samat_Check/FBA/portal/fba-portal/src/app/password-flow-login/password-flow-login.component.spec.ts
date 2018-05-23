import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { PasswordFlowLoginComponent } from './password-flow-login.component';

describe('PasswordFlowLoginComponent', () => {
  let component: PasswordFlowLoginComponent;
  let fixture: ComponentFixture<PasswordFlowLoginComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ PasswordFlowLoginComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(PasswordFlowLoginComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
