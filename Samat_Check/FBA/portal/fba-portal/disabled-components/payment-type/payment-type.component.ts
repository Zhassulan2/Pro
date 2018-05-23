import { Component, OnInit } from '@angular/core';
import { RestService } from '../rest.service';
import { PaymentType } from "../models/paymentType";

@Component({
  selector: 'app-payment-type',
  templateUrl: './payment-type.component.html',
  styleUrls: ['./payment-type.component.css']
})
export class PaymentTypeComponent implements OnInit {
  paymentTypes: Array<PaymentType> = [];
  selectedPaymentType: PaymentType;
  clientSecret: string;
  paymentTypesScope: boolean = false;
  searchScope: boolean = false;
  paginationPaymentTypes: boolean = false;
  paymentTypePages = [];
  paymentTypeSize: number = 10;
  currentPage: number = 1;
  
  constructor(
    private rest: RestService
  ) { }

  ngOnInit() {
    this.getPaymentTypes();
  }

  refreshPaymentType() {
    this.getPaymentTypes();
  }

  onSelect(paymentType: PaymentType): void {
    this.selectedPaymentType = paymentType;
  }

  getPagination() {
    this.rest.getPaymentTypesCount().subscribe(
      res => {
        let paymentTypeCount = parseInt(res.toString());
        if (paymentTypeCount > this.paymentTypeSize) {
          this.paymentTypePages = [];
          for (var i = 0; i < Math.ceil(paymentTypeCount / this.paymentTypeSize); i++) {
            this.paymentTypePages.push(i + 1)
          }
          this.paginationPaymentTypes = true;
        } else {
          this.paymentTypePages = [];
          this.paginationPaymentTypes = false;
        }
      }
    )
  }

  getPaymentTypes(page?: number) {
    this.paymentTypes = [];
    if (page) {
      this.currentPage = page;
    }
    this.rest.getPaymentTypes(page, 10).subscribe(
      (paymentTypes) => {
        this.paymentTypes=paymentTypes;
        this.paymentTypesScope = true;
        this.getPagination();
      }
    );
  }

  createPaymentType(name: string) {
    let postPaymentType: PaymentType = {Name: name};
    this.rest.postPaymentTypes(postPaymentType).subscribe(
      (postPaymentType) => {
        this.getPaymentTypes();
      }
    );
  }

  updatePaymentType() {
    this.rest.putPaymentTypes(this.selectedPaymentType).subscribe(
      (selectedPaymentType) => {
        this.getPaymentTypes();
      }
    );
  }

  deletePaymentType() {
    if (confirm("Вы уверены что хотите удалить точку " + this.selectedPaymentType.Name + "?")) {
      this.rest.deletePaymentTypes(this.selectedPaymentType).subscribe(
        (selectedPaymentType) => {
          this.getPaymentTypes();
        }
      );
    }
  }
}
