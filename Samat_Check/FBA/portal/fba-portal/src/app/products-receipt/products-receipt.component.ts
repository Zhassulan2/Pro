import { Component, OnInit } from '@angular/core';
import { Product } from "../models/product";
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
import { Alerts } from "../alerts.model";

@Component({
  selector: 'app-products-receipt',
  templateUrl: './products-receipt.component.html',
  styleUrls: ['./products-receipt.component.css']
})
export class ProductsReceiptComponent implements OnInit {
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

  //Products
  products: Array<Product> = [];
  paginationProducts: boolean = false;
  productPages = [];
  productSize: number = 10;
  currentPageProduct: number = 1;
  productsScope: boolean = false;
  selectedProduct: Product;


  selectedActionTypeID: string;
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
  lastPrice: number;
  avgPrice: number;

  constructor(
    private rest: RestService,
    private router: Router,
    private alerts: Alerts
  ) { }

  ngOnInit() {
    //this.getProducts();   
    this.getTax();
    this.getActionTypes();
    this.getPaymentTypes();
    this.getSuppliers();
    this.getCompanies();
    this.rest.getActionTypeByShortName("acceptance").subscribe(
      (actionType) => {
        this.selectedActionTypeID=actionType.ID;
      }
    );
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

  //Products
  getProducts(page?: number, name?: string, barcode?: string) {
    this.products = [];
    if (page) {
      this.currentPageProduct = page;
    }
    this.rest.getProducts(page, 10, name, barcode).subscribe(
      (products) => {
        this.products=products;
        if (this.products.length == 1) {
          this.productsScope = false;
          this.addProduct(this.products[0]);
        } else if (this.products.length == 0) {
          this.productsScope = false;
          this.alerts.pushAlert("warning", "ALERTS.SEARCHNOTFOUND");
        }  else {
          this.productsScope = true;
          this.getPaginationProduct();
        }
      }
    );
  }

  getPaginationProduct() {
    this.rest.getProductsCount().subscribe(
      res => {
        let productCount = parseInt(res.toString());
        if (productCount > this.productSize) {
          this.productPages = [];
          for (var i = 0; i < Math.ceil(productCount / this.productSize); i++) {
            this.productPages.push(i + 1)
          }
          this.paginationProducts = true;
        } else {
          this.productPages = [];
          this.paginationProducts = false;
        }
      }
    )
  }

  refreshProduct() {
    this.getProducts();
  }

  backToProducts() {
    this.productActionDetail = null;
  }

  onSelectProduct(product: Product): void {
    this.selectedProduct = product;
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
      }
    );
  }
  
  addProduct(product: Product) {
    if(this.receiptScope.find(x => x.ProductID == product.ID)) {
      this.alerts.pushAlert("danger", "ALERTS.RECEIPTPRODUCT");
    } else {
      this.productActionDetail = {
        ProductID: product.ID,
        Product: product,
        ProductActionID: this.currentProductAction.ID, 
        ActionTypeID: this.selectedActionTypeID,
        Pricebuy: 0,
        Pricesell: 0,
        Count: 1
      };
      this.rest.getProductAvgPrice(product.ID, this.selectedPoint.ID).subscribe(
        res => {
          this.avgPrice = parseInt(res.toString());
        }
      );
      this.rest.getProductLastPrice(product.ID, this.selectedPoint.ID).subscribe(
        res => {
          this.lastPrice = parseInt(res.toString());
        }
      );
    }
  }

  sendProduct() {
    if ((this.productActionDetail.Count <= 0) || (this.productActionDetail.Pricebuy < 0) || (this.productActionDetail.Pricesell < 0)) {
      this.alerts.pushAlert("success", "ALERTS.RECEIPTSEND");
    } else {
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
  }

  deleteScopeMember(scopeMember: ProductActionDetail) {
    let index = this.receiptScope.indexOf(scopeMember,0);
    this.receiptScope.splice(index, 1);
  }

  sendReceiptScope() {
    this.rest.postProductActionDetails(this.receiptScope).subscribe(
      (productActionDetail) => {
        if (productActionDetail) {
          this.router.navigate(["product-action/" + this.currentProductAction.ID]);
        } else {
          this.alerts.pushAlert("danger", "ALERTS.RECEIPTERR");
        }
      }
    );
  }

}
