import { Component, OnInit} from '@angular/core';
import { RestService } from '../rest.service';
import { PermissionCheckService } from '../permission-check.service';
import { Point } from "../models/point";
import { Client } from "../models/client";
import { City } from "../models/city";
import { Company } from "../models/company";

@Component({
  selector: 'app-points',
  templateUrl: './points.component.html',
  styleUrls: ['./points.component.css']
})
export class PointsComponent implements OnInit {
  //New Point
  newName:string;
  newAddress:string;
  newCity:string = "";
  newCompany:string = "";

  points: Array<Point> = [];
  cities: Array<City> = [];
  companies: Array<Company> = [];
  selectedPoint: Point;
  clientSecret: string;
  pointsScope: boolean = false;
  searchScope: boolean = false;
  paginationPoints: boolean = false;
  pointPages = [];
  pointSize: number = 10;
  currentPage: number = 1;
  searchName: string;

  constructor(
    private rest: RestService,
    private check: PermissionCheckService
  ) { }

  ngOnInit() {
    this.getPoints();
    this.getCities();
    this.getCompanies();
  }

  refreshPoint() {
    this.searchName = null;
    this.getPoints();
  }

  onSelect(point: Point): void {
    this.selectedPoint = point;
  }

  getPagination() {
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

  getCities() {
    this.cities = [];
    this.rest.getCitiesCount().subscribe(
      res => {
        let citiesCount = parseInt(res.toString());
        this.rest.getCities(1, citiesCount).subscribe(
          (cities) => {
            this.cities=cities;
          }
        );
      }
    )
  }

  getCompanies() {
    this.companies = [];
    this.rest.getCompaniesCount().subscribe(
      res => {
        let companiesCount = parseInt(res.toString());
        this.rest.getCompanies(1, companiesCount).subscribe(
          (companies) => {
            this.companies=companies;
          }
        );
      }
    )
  }

  getPoints(page?: number) {
    this.points = [];
    if (page) {
      this.currentPage = page;
    }
    this.rest.getPoints(page, 10).subscribe(
      (points) => {
        this.points=points;
        this.pointsScope = true;
        this.getPagination();
      }
    );
  }

  createPoint() {
    let postPoint: Point = {
      Name: this.newName, 
      Address: this.newAddress, 
      CompanyID: this.newCompany
    };
    if (this.newCity != "") {
      postPoint.CityID = this.newCity;
    }
    this.rest.postPoints(postPoint).subscribe(
      (postPoint) => {
        this.getPoints();
        this.newName = null;
        this.newAddress = null;
        this.newCity = "";
        this.newCompany = "";
      }
    );
  }

  updatePoint() {
    this.rest.putPoints(this.selectedPoint).subscribe(
      (selectedPoint) => {
        this.getPoints();
      }
    );
  }

  deletePoint() {
    if (confirm("Вы уверены что хотите удалить точку " + this.selectedPoint.Name + "?")) {
      this.rest.deletePoints(this.selectedPoint).subscribe(
        (selectedPoint) => {
          this.getPoints();
        }
      );
    }
  }

  getClientSecret() {
    let client: Client = {ID:this.selectedPoint.ID}
    this.rest.getClientInfo(client).subscribe(
      (client) => {
        this.clientSecret = client.Secret;
      }
    );
  }
}