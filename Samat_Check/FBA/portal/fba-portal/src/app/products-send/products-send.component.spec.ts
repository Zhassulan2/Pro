import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ProductsSendComponent } from './products-send.component';
import {beforeEach, describe, expect, it} from '@angular/core/testing/src/testing_internal';

describe('ProductsSendComponent', () => {
  let component: ProductsSendComponent;
  let fixture: ComponentFixture<ProductsSendComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ProductsSendComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ProductsSendComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
