import { Component, OnInit } from '@angular/core';
import { Tax } from "../models/tax";
import { Company } from "../models/company";
import { Point } from "../models/point";
import { Supplier } from "../models/supplier";
import { ActionType } from "../models/actionType";
import { PaymentType } from "../models/paymentType";
import { ProductActionDetail } from "../models/productActionDetail";
import { ProductAction } from "../models/productAction";
import { RestService } from '../rest.service';
import { ProductStock } from "../models/productStock";
import { Router } from '@angular/router';

@Component({
  selector: 'app-products-send',
  templateUrl: './products-send.component.html',
  styleUrls: ['./products-send.component.css']
})
export class ProductsSendComponent implements OnInit {
  //Point
  currentPagePoint: number = 1;
  selectedPoint: Point;
  paginationPoints: boolean = false;
  pointPages = [];
  pointSize: number = 10;
  points: Array<Point> = [];

  //Company
  currentPageCompany: number = 1;
  companies: Array<Company> = [];
  selectedCompany: Company;
  paginationCompanies: boolean = false;
  companyPages = [];
  companySize: number = 10;


  //ProductStock
  productStock: Array<ProductStock> = [];
  paginationProductStock: boolean = false;
  productStockPages = [];
  productStockSize: number = 10;
  currentPageProductStock: number = 1;
  productStockScope: boolean = false;
  selectedProductStock: ProductStock;


  selectedActionTypeID: string = 'e2e935ca-eb13-47c5-95e9-3c15923d77ef';
  taxes: Array<Tax> = [];
  suppliers: Array<Supplier> = [];
  actionTypes: Array<ActionType> = [];
  paymentTypes: Array<PaymentType> = [];
  productActionDetail: ProductActionDetail;
  receiptScope: Array<ProductActionDetail> = [];
  extraFee: number = 0;
  extraFeeValue: number = 0;
  extraFeePercent: number = 0;
  currentProductAction: ProductAction;
  pointA: Point;
  pointB: Point;

  constructor(
    private rest: RestService,
    private router: Router
  ) { }

  ngOnInit() {
    this.getCompanies();
  }

  //Points
  getPoints(page?: number) {
    this.points = [];
    if (page) {
      this.currentPagePoint = page;
    }
    this.rest.getPoints(page, 10, this.selectedCompany.ID).subscribe(
      (points) => {
        this.points=points;
        this.getPaginationPoints();
      }
    );
  }

  getPaginationPoints() {
    this.rest.getPointsCount().subscribe(
      res => {
        let pointCount = parseInt(res.toString());
        if (pointCount > this.pointSize) {
          this.pointPages = [];
          for (var i = 0; i < Math.ceil(pointCount / this.pointSize); i++) {
            this.pointPages.push(i + 1)
          }
          this.paginationPoints = true;
        } else {
          this.pointPages = [];
          this.paginationPoints = false;
        }
      }
    )
  }

  onSelectPoint(point: Point): void {
    this.selectedPoint = point;
  }

  //Companies
  getCompanies(page?: number) {
    this.companies = [];
    if (page) {
      this.currentPageCompany = page;
    }
    this.rest.getCompanies(page, 10).subscribe(
      (companies) => {
        this.companies=companies;
        this.getPaginationCompany();
      }
    );
  }

  getPaginationCompany() {
    this.rest.getCompaniesCount().subscribe(
      res => {
        let companyCount = parseInt(res.toString());
        if (companyCount > this.companySize) {
          this.companyPages = [];
          for (var i = 0; i < Math.ceil(companyCount / this.companySize); i++) {
            this.companyPages.push(i + 1)
          }
          this.paginationCompanies = true;
        } else {
          this.companyPages = [];
          this.paginationCompanies = false;
        }
      }
    )
  }

  onSelectCompany(company: Company): void {
    this.selectedCompany = company;
    this.getPoints();
  }

  //ProductStock
  getProductStock(page?: number, name?: string, barcode?: string) {
    this.productStock = [];
    if (page) {
      this.currentPageProductStock = page;
    }
    this.rest.getProductStock(page, 10).subscribe(
      (productStock) => {
        this.productStock=productStock;
        this.productStockScope = true;
        this.getPaginationProductStock();
      }
    );
  }

  getPaginationProductStock() {
    this.rest.getProductsCount().subscribe(
      res => {
        let productCount = parseInt(res.toString());
        if (productCount > this.productStockSize) {
          this.productStockPages = [];
          for (var i = 0; i < Math.ceil(productCount / this.productStockSize); i++) {
            this.productStockPages.push(i + 1)
          }
          this.paginationProductStock = true;
        } else {
          this.productStockPages = [];
          this.paginationProductStock = false;
        }
      }
    )
  }

  refreshProductStock() {
    this.getProductStock();
  }

  backToProductStock() {
    this.productActionDetail = null;
  }

  onSelectProductStock(productStock: ProductStock): void {
    this.selectedProductStock = productStock;
  }


  //Dictionaries
  getTax() {
    this.taxes = [];
    this.rest.getTaxesCount().subscribe(
      res => {
        let count = parseInt(res.toString());
        this.rest.getTaxes(1, count).subscribe(
          (res) => {
            this.taxes=res;
          }
        );
      }
    )
  }

  getActionTypes() {
    this.actionTypes = [];
    this.rest.getActionTypesCount().subscribe(
      res => {
        let count = parseInt(res.toString());
        this.rest.getActionTypes(1, count).subscribe(
          (res) => {
            this.actionTypes=res;
          }
        );
      }
    )
  }

  getPaymentTypes() {
    this.paymentTypes = [];
    this.rest.getPaymentTypesCount().subscribe(
      res => {
        let count = parseInt(res.toString());
        this.rest.getPaymentTypes(1, count).subscribe(
          (res) => {
            this.paymentTypes=res;
          }
        );
      }
    )
  }

  getSuppliers() {
    this.suppliers = [];
    this.rest.getSuppliersCount().subscribe(
      res => {
        let count = parseInt(res.toString());
        this.rest.getSuppliers(1, count).subscribe(
          (res) => {
            this.suppliers=res;
          }
        );
      }
    )
  }

  //Receipt logic
  setPointA() {
    if (this.pointB && this.pointB == this.selectedPoint) {
      alert("Точка отпраки и точка получения не могут совпадать!")
    } else {
      this.pointA = this.selectedPoint;
    }
  }

  setPointB() {
    if (this.pointA && this.pointA == this.selectedPoint) {
      alert("Точка отпраки и точка получения не могут совпадать!")
    } else {
      this.pointB = this.selectedPoint;
    }
  }


  setExtraFee() {
    if (this.extraFee == 1) {
      this.productActionDetail.Pricesell = this.productActionDetail.Pricebuy + this.extraFeeValue;
    } else if (this.extraFee == 2) {
      this.productActionDetail.Pricesell = this.productActionDetail.Pricebuy + ((this.productActionDetail.Pricebuy * this.extraFeePercent)/100);
    } else {
      this.productActionDetail.Pricesell = this.productActionDetail.Pricebuy;
    }
  }

  startAction() {
    this.receiptScope = [];
    let currentDate: Date = new Date();
    let productAction: ProductAction = {
      PointID: this.selectedPoint.ID,
      ActionDate: currentDate.toISOString(),
      Count: 0,
      Amount: 0
    };
    this.rest.postProductAction(productAction).subscribe(
      (productAction) => {
        this.currentProductAction = productAction;
        this.getProductStock();
      }
    );
  }
  
  addProductStock(productStock: ProductStock) {
    if(this.receiptScope.find(x => x.ProductID == productStock.ProductID)) {
      alert("Данный продукт уже есть в списке на приёмку");
    } else {
      this.productActionDetail = {
        ProductID: productStock.ProductID,
        Product: productStock.Product,
        ProductActionID: this.currentProductAction.ID, 
        ActionTypeID: this.selectedActionTypeID,
        Pricebuy: 0,
        Pricesell: 0,
        Count: 1
      };
    }
  }

  sendProductStock() {
    if (this.productActionDetail.TaxID) {
      this.productActionDetail.Tax = this.taxes.find(x => x.ID == this.productActionDetail.TaxID);
    }
    if (this.productActionDetail.SupplierID) {
      this.productActionDetail.Supplier = this.suppliers.find(x => x.ID == this.productActionDetail.SupplierID);
    }
    if (this.productActionDetail.PaymentTypeID) {
      this.productActionDetail.PaymentType = this.paymentTypes.find(x => x.ID == this.productActionDetail.PaymentTypeID);
    }
    this.receiptScope.push(this.productActionDetail);
    this.productActionDetail = null;
  }

  deleteScopeMember(scopeMember: ProductActionDetail) {
    let index = this.receiptScope.indexOf(scopeMember,0);
    this.receiptScope.splice(index, 1);
  }

  sendReceiptScope() {
    for (let receiptMember of this.receiptScope) {
      this.rest.postProductActionDetail(receiptMember);
      let productStock: ProductStock = {
        ProductID: receiptMember.ProductID,
        PointID: this.selectedPoint.ID,
        Count: receiptMember.Count,
        PriceSell: receiptMember.Pricesell
      }
      this.rest.postProductStock(productStock);
    }
    this.router.navigate(["product-action/" + this.currentProductAction.ID]);
  }

}
