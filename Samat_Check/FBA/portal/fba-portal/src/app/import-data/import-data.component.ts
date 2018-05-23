import { Component, OnInit } from '@angular/core';
import { Alerts } from "../alerts.model";
import { ImportProduct } from "../models/importProduct";
import { Company } from "../models/company";
import { Point } from "../models/point";
import { RestService } from '../rest.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-import-data',
  templateUrl: './import-data.component.html',
  styleUrls: ['./import-data.component.css']
})
export class ImportDataComponent implements OnInit {
  fileToUpload: File = null;
  fileInfo: string = "...";
  importProduct: Array<ImportProduct> = [];
  importProductPreview: Array<ImportProduct> = [];
  paginationPreview: boolean = false;
  previewPages = [];
  previewSize: number = 10;
  currentPreviewPage: number = 1;
  pageCount: number = 0;

  startPagination: number = 0;
  endPagination: number = 10;

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

  constructor(private rest: RestService, private router: Router, private alerts: Alerts) { }

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

  //Import logic
  getPreview(page?: number) {
    this.importProductPreview = [];
    if (page) {
      this.currentPreviewPage = page;
    }
    let sliceStart: number = ((this.currentPreviewPage-1) * this.previewSize);
    let sliceEnd: number = sliceStart + this.previewSize;
    this.importProductPreview = this.importProduct.slice(sliceStart, sliceEnd);
    this.getPaginationPreview();
  }

  getPaginationPreview() {
    if (this.importProduct.length > this.previewSize) {
      this.previewPages = [];
      this.pageCount = Math.ceil(this.importProduct.length / this.previewSize);
      if (this.startPagination - this.currentPreviewPage < -5) {
        this.startPagination = this.startPagination + 5; //i
        this.endPagination = this.startPagination + 10;
      } else if ((this.startPagination - this.currentPreviewPage >= 0)){
        this.startPagination = this.startPagination - 5; //i
        this.endPagination = this.startPagination + 10;
      }
      for (var i = this.startPagination; (i < this.pageCount && (i < this.endPagination)); i++) {
        this.previewPages.push(i + 1)
      }
      this.paginationPreview = true;
    } else {
      this.previewPages = [];
      this.paginationPreview = false;
    }
  }


  handleFileInput(files: FileList) {
    if (files.item(0).type == "text/csv") {
      this.importProduct = [];
      this.fileToUpload = files.item(0);
      this.fileInfo = this.fileToUpload.name;

      let reader: FileReader = new FileReader();
      reader.readAsText(this.fileToUpload);
      reader.onload = (e) => {
        let csv: string = reader.result;
        let lines = csv.split(/\r\n|\r|\n/);
        for ( let i = 1; i < lines.length; i++) {
          let importProductActionRaw: ImportProduct = {
            Name: lines[i].split(";")[0],
            BarCode: lines[i].split(";")[1],
            PriceBuy: Number(lines[i].split(";")[2]),
            PriceSell: Number(lines[i].split(";")[3]),
            Count: Number(lines[i].split(";")[4]),
            CategoryName: lines[i].split(";")[5]
          }
          this.importProduct.push(importProductActionRaw);
        }
        this.getPreview();
      }
    } else {
      this.alerts.pushAlert("danger", "ALERTS.FILEFORMATERR");
    }
  }

  rejectImport() {
    this.importProduct = [];
    this.fileToUpload = null;
    this.fileInfo = "...";
    this.selectedPoint = null;
    this.selectedCompany = null;
    this.startPagination = 0;
    this.endPagination = 10;
  }

  startImport() {
    this.rest.postImportProduct(this.importProduct, this.selectedPoint.ID).subscribe(
      (productAction) => {
        if (productAction) {
          this.router.navigate(["product-action/" + productAction.ID]);
        } else {
          this.alerts.pushAlert("danger", "ALERTS.RECEIPTERR");
        }
      }
    );
  }
}
