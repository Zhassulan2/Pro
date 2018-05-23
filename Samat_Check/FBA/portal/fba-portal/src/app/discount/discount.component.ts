import { Component, OnInit } from '@angular/core';
import { RestService } from '../rest.service';
import { Discount } from "../models/discount";
import { DiscountType } from "../models/discountType";

@Component({
  selector: 'app-discount',
  templateUrl: './discount.component.html',
  styleUrls: ['./discount.component.css']
})
export class DiscountComponent implements OnInit {
  //New Discount
  newName: string;
  newDiscountTypeID: string = "";
  newStartDate: Date;
  newEndDate: Date;
  newValue: number;
  newValuepercent: number;

  //Put Discount
  putStart: Date;
  putEnd: Date;

  discountTypes: Array<DiscountType> = [];
  discounts: Array<Discount> = [];
  selectedDiscount: Discount;
  discountsScope: boolean = false;
  searchScope: boolean = false;
  paginationDiscounts: boolean = false;
  discountPages = [];
  discountSize: number = 10;
  currentPage: number = 1;
  
  constructor(
    private rest: RestService
  ) { }

  ngOnInit() {
    this.getDiscounts();
    this.getDiscountTypes();
  }

  getDiscountTypes() {
    this.discountTypes = [];
    this.rest.getDiscountsCount().subscribe(
      res => {
        let discountsCount = parseInt(res.toString());
        this.rest.getDiscountTypes(1, discountsCount).subscribe(
          (discountTypes) => {
            this.discountTypes=discountTypes;
          }
        );
      }
    )
  }

  refreshDiscount() {
    this.getDiscounts();
  }

  onSelect(discount: Discount): void {
    this.selectedDiscount = discount;
    this.putStart = new Date(this.selectedDiscount.Startdate);
    this.putEnd = new Date(this.selectedDiscount.Enddate);
  }

  getPagination() {
    this.rest.getDiscountsCount().subscribe(
      res => {
        let discountCount = parseInt(res.toString());
        if (discountCount > this.discountSize) {
          this.discountPages = [];
          for (var i = 0; i < Math.ceil(discountCount / this.discountSize); i++) {
            this.discountPages.push(i + 1)
          }
          this.paginationDiscounts = true;
        } else {
          this.discountPages = [];
          this.paginationDiscounts = false;
        }
      }
    )
  }

  getDiscounts(page?: number) {
    this.discounts = [];
    if (page) {
      this.currentPage = page;
    }
    this.rest.getDiscounts(page, 10).subscribe(
      (discounts) => {
        this.discounts=discounts;
        this.discountsScope = true;
        this.getPagination();
      }
    );
  }

  createDiscount() {
    let postDiscount: Discount = {
      Name: this.newName,
      DiscountTypeID: this.newDiscountTypeID,
      Startdate: this.newStartDate.toISOString(),
      Enddate: this.newEndDate.toISOString(),
      Value: this.newValue,
      Valuepercent: this.newValuepercent
    };
    this.rest.postDiscounts(postDiscount).subscribe(
      (postDiscount) => {
        this.getDiscounts();
        this.newName = null;
        this.newDiscountTypeID = "";
        this.newStartDate = null;
        this.newEndDate = null;
        this.newValue = null;
        this.newValuepercent = null;
      }
    );
  }

  updateDiscount() {
    this.selectedDiscount.Startdate = this.putStart.toISOString();
    this.selectedDiscount.Enddate = this.putEnd.toISOString();
    this.rest.putDiscounts(this.selectedDiscount).subscribe(
      (selectedDiscount) => {
        this.getDiscounts();
      }
    );
  }

  deleteDiscount() {
    if (confirm("Вы уверены что хотите удалить точку " + this.selectedDiscount.Name + "?")) {
      this.rest.deleteDiscounts(this.selectedDiscount).subscribe(
        (selectedDiscount) => {
          this.getDiscounts();
        }
      );
    }
  }
}
