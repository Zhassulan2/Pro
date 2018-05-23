import { Component, OnInit } from '@angular/core';
import { NgbModule, NgbDateStruct, NgbDateParserFormatter } from '@ng-bootstrap/ng-bootstrap';
import { RestService } from '../rest.service';
import { Tax } from "../models/tax";
import { toInteger } from '@ng-bootstrap/ng-bootstrap/util/util';

@Component({
  selector: 'app-tax',
  templateUrl: './tax.component.html',
  styleUrls: ['./tax.component.css']
})
export class TaxComponent implements OnInit {
  taxes: Array<Tax> = [];
  selectedTax: Tax;
  clientSecret: string;
  taxesScope: boolean = false;
  searchScope: boolean = false;
  paginationTaxes: boolean = false;
  taxPages = [];
  taxSize: number = 10;
  currentPage: number = 1;
  validFrom: NgbDateStruct;
  validFromSelected: NgbDateStruct;

  constructor(
    private rest: RestService,
    private parserFormatter: NgbDateParserFormatter
  ) { }

  ngOnInit() {
    this.getTaxes();
  }

  refreshTax() {
    this.getTaxes();
  }

  onSelect(tax: Tax): void {
    this.selectedTax = tax;
    let tempDate: Date = new Date(this.selectedTax.ValidFrom);
    this.validFromSelected = {year: tempDate.getFullYear(), month: tempDate.getMonth() + 1, day: tempDate.getDate()};
  }

  getPagination() {
    this.rest.getTaxesCount().subscribe(
      res => {
        let taxCount = parseInt(res.toString());
        if (taxCount > this.taxSize) {
          this.taxPages = [];
          for (var i = 0; i < Math.ceil(taxCount / this.taxSize); i++) {
            this.taxPages.push(i + 1)
          }
          this.paginationTaxes = true;
        } else {
          this.taxPages = [];
          this.paginationTaxes = false;
        }
      }
    )
  }

  getTaxes(page?: number) {
    this.taxes = [];
    if (page) {
      this.currentPage = page;
    }
    this.rest.getTaxes(page, 10).subscribe(
      (taxes) => {
        this.taxes=taxes;
        this.taxesScope = true;
        this.getPagination();
      }
    );
  }

  createTax(name: string, validFrom: NgbDateStruct, taxValue: number) {
    let tempValidFrom = this.parserFormatter.format(validFrom) + "T00:00:00Z";
    let postTax: Tax = {Name: name, ValidFrom:tempValidFrom, Value: Number(taxValue)};
    this.rest.postTaxes(postTax).subscribe(
      (postTax) => {
        this.getTaxes();
      }
    );
  }

  updateTax() {
    let tempValidFrom = this.parserFormatter.format(this.validFromSelected) + "T00:00:00Z";
    if (tempValidFrom != this.selectedTax.ValidFrom) {
      this.selectedTax.ValidFrom = tempValidFrom;
      this.validFromSelected = null;
    }
    this.rest.putTaxes(this.selectedTax).subscribe(
      (selectedTax) => {
        this.getTaxes();
      }
    );
  }

  deleteTax() {
    if (confirm("Вы уверены что хотите удалить точку " + this.selectedTax.Name + "?")) {
      this.rest.deleteTaxes(this.selectedTax).subscribe(
        (selectedTax) => {
          this.getTaxes();
        }
      );
    }
  }
}

