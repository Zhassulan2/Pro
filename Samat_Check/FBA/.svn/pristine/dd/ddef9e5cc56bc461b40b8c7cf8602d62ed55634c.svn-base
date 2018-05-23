import { Component, OnInit } from '@angular/core';
import { RestService } from '../rest.service';
import { Inventory } from "../models/inventory";
import { InventoryDetail } from "../models/inventoryDetail";
import { Company } from "../models/company";
import { Point } from "../models/point";
import { Product } from "../models/product";
import { Alerts } from "../alerts.model";

@Component({
  selector: 'app-inventory',
  templateUrl: './inventory.component.html',
  styleUrls: ['./inventory.component.css']
})
export class InventoryComponent implements OnInit {
  //New Inventory
  newName: string;
  newCompanyID: string = "";
  newStartDate: Date = new Date();
  //newEndDate: Date;
  newPointID: string = "";

  //Put Inventory
  putStart: Date;
  putEnd: Date;

  //Products
  products: Array<Product> = [];
  paginationProducts: boolean = false;
  productPages = [];
  productSize: number = 10;
  currentPageProduct: number = 1;
  productsScope: boolean = false;
  selectedProduct: Product;

  //InventoryDetail
  currentPageInvDet: number = 1;
  selectedInvDet: InventoryDetail;
  paginationInvDet: boolean = false;
  inventoryDetailsPages = [];
  inventoryDetailsSize: number = 10;
  inventoryDetails: Array<InventoryDetail> = [];

  inventoryDetail: InventoryDetail;
  points: Array<Point> = [];
  companies: Array<Company> = [];
  inventories: Array<Inventory> = [];
  selectedInventory: Inventory;
  inventoryInWork: Inventory;
  inventoriesScope: boolean = false;
  searchScope: boolean = false;
  paginationInventories: boolean = false;
  inventoryPages = [];
  inventorySize: number = 10;
  currentPage: number = 1;
  
  constructor(
    private rest: RestService,
    private alerts: Alerts
  ) { }

  ngOnInit() {
    this.getInventories();
    this.getCompanies();
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
          this.alerts.pushAlert( "warning", "ALERTS.SEARCHNOTFOUND");
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
    this.inventoryDetail = null;
  }

  onSelectProduct(product: Product): void {
    this.selectedProduct = product;
    this.addProduct(product);
  }

  //InventoryDetail
  getInventoryDetails(page?: number) {
    this.inventoryDetails = [];
    if (page) {
      this.currentPageInvDet = page;
    }
    this.rest.getInventoryDetails(this.inventoryInWork.ID, page, 10).subscribe(
      (inventoryDetails) => {
        this.inventoryDetails=inventoryDetails;
        this.getPaginationInvDet();
      }
    );
  }

  getPaginationInvDet() {
    this.rest.getPointsCount().subscribe(
      res => {
        let invCount = parseInt(res.toString());
        if (invCount > this.inventoryDetailsSize) {
          this.inventoryDetailsPages = [];
          for (var i = 0; i < Math.ceil(invCount / this.inventoryDetailsSize); i++) {
            this.inventoryDetailsPages.push(i + 1)
          }
          this.paginationInvDet = true;
        } else {
          this.inventoryDetailsPages = [];
          this.paginationInvDet = false;
        }
      }
    )
  }

  onSelectInvDet(invDet: InventoryDetail): void {
    this.selectedInvDet = invDet;
  }

  getCompanies() {
    this.companies = [];
    this.rest.getInventoriesCount().subscribe(
      res => {
        let inventoriesCount = parseInt(res.toString());
        this.rest.getCompanies(1, inventoriesCount).subscribe(
          (companies) => {
            this.companies=companies;
          }
        );
      }
    )
  }

  refreshInventory() {
    this.getInventories();
  }

  getPoints(companyID: string) {
    this.points = [];
    this.rest.getPointsCount().subscribe(
      res => {
        let pointsCount = parseInt(res.toString());
        this.rest.getPoints(1, pointsCount, companyID).subscribe(
          (points) => {
            this.points=points;
          }
        );
      }
    )
  }

  onSelect(inventory: Inventory): void {
    this.selectedInventory = inventory;
    this.putStart = new Date(this.selectedInventory.Startdate);
    this.putEnd = new Date(this.selectedInventory.Enddate);
  }

  getPagination() {
    this.rest.getInventoriesCount().subscribe(
      res => {
        let inventoryCount = parseInt(res.toString());
        if (inventoryCount > this.inventorySize) {
          this.inventoryPages = [];
          for (var i = 0; i < Math.ceil(inventoryCount / this.inventorySize); i++) {
            this.inventoryPages.push(i + 1)
          }
          this.paginationInventories = true;
        } else {
          this.inventoryPages = [];
          this.paginationInventories = false;
        }
      }
    )
  }

  getInventories(page?: number) {
    this.inventories = [];
    if (page) {
      this.currentPage = page;
    }
    this.rest.getInventories(page, 10).subscribe(
      (inventories) => {
        this.inventories=inventories;
        this.inventoriesScope = true;
        this.getPagination();
      }
    );
  }

  showDetails() {
    this.inventoryInWork=this.selectedInventory; 
    this.products = [];
    this.selectedProduct = null;
    this.getInventoryDetails();
  }

  createInventory() {
    let postInventory: Inventory = {
      Name: this.newName,
      CompanyID: this.newCompanyID,
      Startdate: this.newStartDate.toISOString(),
      //Enddate: this.newEndDate.toISOString(),
      PointID: this.newPointID
    };
    this.rest.postInventories(postInventory).subscribe(
      (postInventory) => {
        this.getInventories();
        this.newName = null;
        this.newCompanyID = "";
        this.newStartDate = new Date();
        //this.newEndDate = null;
        this.newPointID = "";
      }
    );
  }

  endInventory() {
    let currentDate: Date = new Date();
    this.selectedInventory.Enddate = currentDate.toISOString();
    this.rest.putInventories(this.selectedInventory).subscribe(
      (selectedInventory) => {
        this.getInventories();
        this.inventoryInWork = null;
      }
    );
  }

  updateInventory() {
    this.selectedInventory.Startdate = this.putStart.toISOString();
    this.selectedInventory.Enddate = this.putEnd.toISOString();
    this.rest.putInventories(this.selectedInventory).subscribe(
      (selectedInventory) => {
        this.getInventories();
      }
    );
  }

  deleteInventory() {
    if (confirm("Вы уверены что хотите удалить точку " + this.selectedInventory.Name + "?")) {
      this.rest.deleteInventories(this.selectedInventory).subscribe(
        (selectedInventory) => {
          this.getInventories();
        }
      );
    }
  }

  addProduct(product: Product) {
    if(this.inventoryDetails.find(x => x.ProductID == product.ID)) {
      this.alerts.pushAlert( "danger", "ALERTS.RECEIPTPRODUCT");
    } else {
      this.inventoryDetail = {
        ProductID: product.ID,
        Product: product,
        InventoryID: this.inventoryInWork.ID,
        Count: 0
      };
    }
  }

  sendInventoryDetail() {
    if (this.inventoryDetail.Count > 0) {
      let postInventoryDetail: InventoryDetail = {
        InventoryID: this.inventoryInWork.ID,
        Product: this.inventoryDetail.Product,
        ProductID: this.inventoryDetail.ProductID,
        Count: this.inventoryDetail.Count
      };
      this.rest.postInventoryDetails(postInventoryDetail).subscribe(
        (postInventoryDetail) => {
          this.getInventoryDetails();
          this.backToProducts();
        }
      );
    } else {
      this.alerts.pushAlert( "warning", "ALERTS.INVENTORYSEND");
    }
  }

  deleteInventoryDetail() {
    if (confirm("Вы уверены что хотите удалить точку " + this.selectedProduct.Name + "?")) {
      this.rest.deleteInventoryDetails(this.inventoryDetail).subscribe(
        (inventoryDetail) => {
          this.getInventoryDetails();
        }
      );
    }
  }
}

