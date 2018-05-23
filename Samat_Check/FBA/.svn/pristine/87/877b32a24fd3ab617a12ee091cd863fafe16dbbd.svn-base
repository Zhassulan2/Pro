import { Component, OnInit } from '@angular/core';
import { Company } from "../models/company";
import { Point } from "../models/point";
import { ProductActionDetail } from "../models/productActionDetail";
import { ProductAction } from "../models/productAction";
import { RestService } from '../rest.service';
import { ProductStock } from "../models/productStock";
import { Router } from '@angular/router';
import { ActionType } from '../models/actionType';
import { Alerts } from "../alerts.model";
@Component({
  selector: 'app-product-refuse',
  templateUrl: './product-refuse.component.html',
  styleUrls: ['./product-refuse.component.css']
})
export class ProductRefuseComponent implements OnInit {
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
  actionType: ActionType;

  actionTypes: Array<ActionType> = [];
  selectedTaxID: string;
  selectedPaymentTypeID: string;
  selectedActionTypeID: string;
  productActionDetail: ProductActionDetail;
  refuseScope: Array<ProductActionDetail> = [];
  currentProductAction: ProductAction;


  constructor(
    private rest: RestService,
    private router: Router,
    private alerts: Alerts
  ) { }

  ngOnInit() {
    this.getCompanies();
    this.getActionTypes();
    
    this.rest.getTaxByShortName("notax").subscribe(
      (tax) => {
        this.selectedTaxID=tax.ID;
      }
    );
    this.rest.getActionTypeByShortName("refuse").subscribe(
      (actionType) => {
        this.selectedActionTypeID=actionType.ID;
      }
    );
    this.rest.getPaymentTypeByShortName("nopay").subscribe(
      (paymentType) => {
        this.selectedPaymentTypeID=paymentType.ID;
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

  //Dictionaries
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

  //ProductStock
  getProductStock(page?: number, name?: string, barcode?: string) {
    this.productStock = [];
    if (page) {
      this.currentPageProductStock = page;
    }
    this.rest.getProductStockByPID(this.selectedPoint.ID, page, 10, name, barcode).subscribe(
      (productStock) => {
        this.productStock=productStock;
        if (this.productStock.length == 1) {
          this.selectedProductStock = this.productStock[0];
          this.productStockScope = false;
          this.addProductStock();
        } else if (this.productStock.length == 0) {
          this.productStockScope = false;
          this.alerts.pushAlert("warning", "ALERTS.SEARCHNOTFOUND");
        } else {
          this.productStockScope = true;
          this.getPaginationProductStock();
        }
      }
    );
  }

  getPaginationProductStock() {
    this.rest.getProductStockCount().subscribe(
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

  //Send logic
  startAction() {
    this.refuseScope = [];
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
  
  addProductStock() {
    if(this.refuseScope.find(x => x.ProductID == this.selectedProductStock.ProductID)) {
      this.alerts.pushAlert("danger", "ALERTS.RELOCATIONPRODUCT");
    } else {
      this.productActionDetail = {
        ProductID: this.selectedProductStock.ProductID,
        Product: this.selectedProductStock.Product,
        TaxID: this.selectedTaxID,
        PaymentTypeID: this.selectedPaymentTypeID,
        ProductActionID: this.currentProductAction.ID, 
        ActionTypeID: this.selectedActionTypeID,
        Pricesell: this.selectedProductStock.PriceSell,
        Count: 0
      };
    }
  }

  sendProductStock() {
    if ((this.productActionDetail.Count > this.selectedProductStock.Count) || (this.productActionDetail.Count < 0)) {
      this.alerts.pushAlert("danger", "ALERTS.RELOCATIONSEND");
    } else {      
      this.productActionDetail.Count = this.productActionDetail.Count * -1;
      this.refuseScope.push(this.productActionDetail);
      this.productActionDetail = null;
    }
  }

  deleteScopeMember(scopeMember: ProductActionDetail) {
    let index = this.refuseScope.indexOf(scopeMember,0);
    this.refuseScope.splice(index, 1);
  }

  sendRefuseScope() {
    this.rest.postProductActionDetails(this.refuseScope).subscribe(
      (productActionDetail) => {
        if (productActionDetail) {
          this.router.navigate(["product-action/" + this.currentProductAction.ID]);
        } else {
          this.alerts.pushAlert("danger", "ALERTS.RELOCATIONERR");
        }
      }
    );
  }
}
