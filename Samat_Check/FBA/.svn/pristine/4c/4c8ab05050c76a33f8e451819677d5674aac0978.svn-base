import { Component, OnInit } from '@angular/core';
import { RestService } from '../rest.service';
import { Supplier } from "../models/supplier";
import { PermissionCheckService } from '../permission-check.service';

@Component({
  selector: 'app-supplier',
  templateUrl: './supplier.component.html',
  styleUrls: ['./supplier.component.css']
})
export class SupplierComponent implements OnInit {
  //New Supplier
  newName:string;
  newAddress:string;
  newBIN:string;
  newContact:string;

  suppliers: Array<Supplier> = [];
  selectedSupplier: Supplier;
  clientSecret: string;
  suppliersScope: boolean = false;
  searchScope: boolean = false;
  paginationSuppliers: boolean = false;
  supplierPages = [];
  supplierSize: number = 10;
  currentPage: number = 1;
  
  constructor(
    private rest: RestService,
    private check: PermissionCheckService
  ) { }

  ngOnInit() {
    this.getSuppliers();
  }

  refreshSupplier() {
    this.getSuppliers();
  }

  onSelect(supplier: Supplier): void {
    this.selectedSupplier = supplier;
  }

  getPagination() {
    this.rest.getSuppliersCount().subscribe(
      res => {
        let supplierCount = parseInt(res.toString());
        if (supplierCount > this.supplierSize) {
          this.supplierPages = [];
          for (var i = 0; i < Math.ceil(supplierCount / this.supplierSize); i++) {
            this.supplierPages.push(i + 1)
          }
          this.paginationSuppliers = true;
        } else {
          this.supplierPages = [];
          this.paginationSuppliers = false;
        }
      }
    )
  }

  getSuppliers(page?: number) {
    this.suppliers = [];
    if (page) {
      this.currentPage = page;
    }
    this.rest.getSuppliers(page, 10).subscribe(
      (suppliers) => {
        this.suppliers=suppliers;
        this.suppliersScope = true;
        this.getPagination();
      }
    );
  }

  createSupplier() {
    let postSupplier: Supplier = {Name: this.newName, Address: this.newAddress, Bin: this.newBIN, Contact: this.newContact};
    this.rest.postSuppliers(postSupplier).subscribe(
      (postSupplier) => {
        this.getSuppliers();
        this.newName = null;
        this.newAddress = null;
        this.newBIN = null;
        this.newContact = null;
      }
    );
  }

  updateSupplier() {
    this.rest.putSuppliers(this.selectedSupplier).subscribe(
      (selectedSupplier) => {
        this.getSuppliers();
      }
    );
  }

  deleteSupplier() {
    if (confirm("Вы уверены что хотите удалить точку " + this.selectedSupplier.Name + "?")) {
      this.rest.deleteSuppliers(this.selectedSupplier).subscribe(
        (selectedSupplier) => {
          this.getSuppliers();
        }
      );
    }
  }
}
