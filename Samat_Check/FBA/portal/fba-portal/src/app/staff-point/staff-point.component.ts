import { Component, OnInit } from '@angular/core';
import { Point } from "../models/point";
import { Staff } from "../models/staff";
import { StaffPoint } from "../models/staffPoint";
import { RestService } from '../rest.service';
import { Alerts } from "../alerts.model";

@Component({
  selector: 'app-staff-point',
  templateUrl: './staff-point.component.html',
  styleUrls: ['./staff-point.component.css']
})
export class StaffPointComponent implements OnInit {
  //Point
  currentPagePoint: number = 1;
  selectedPoint: Point;
  paginationPoints: boolean = false;
  pointPages = [];
  pointSize: number = 10;
  points: Array<Point> = [];

  //Staff
  currentPageStaff: number = 1;
  staffs: Array<Staff> = [];
  selectedStaff: Staff;
  paginationStaffs: boolean = false;
  staffPages = [];
  staffSize: number = 10;

  //StaffPoint
  currentPageSP: number = 1;
  selectedSP: StaffPoint;
  paginationSP: boolean = false;
  spPages = [];
  spSize: number = 10;
  staffPoints: Array<StaffPoint> = [];

  constructor(
    private rest: RestService,
    private alerts: Alerts
  ) { }

  ngOnInit() {
    this.getPoints();
    this.getStaffs();
    this.getStaffPoints();
  }

  //Points
  getPoints(page?: number) {
    this.points = [];
    if (page) {
      this.currentPagePoint = page;
    }
    this.rest.getPoints(page, 10).subscribe(
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

  //Staffs
  getStaffs(page?: number) {
    this.staffs = [];
    if (page) {
      this.currentPageStaff = page;
    }
    this.rest.getStaffs(page, 10).subscribe(
      (staffs) => {
        this.staffs=staffs;
        this.getPaginationStaff();
      }
    );
  }

  getPaginationStaff() {
    this.rest.getStaffsCount().subscribe(
      res => {
        let staffCount = parseInt(res.toString());
        if (staffCount > this.staffSize) {
          this.staffPages = [];
          for (var i = 0; i < Math.ceil(staffCount / this.staffSize); i++) {
            this.staffPages.push(i + 1)
          }
          this.paginationStaffs = true;
        } else {
          this.staffPages = [];
          this.paginationStaffs = false;
        }
      }
    )
  }

  onSelectStaff(staff: Staff): void {
    this.selectedStaff = staff;
  }

  //StaffPoint
  getStaffPoints(page?: number) {
    this.staffPoints = [];
    if (page) {
      this.currentPageSP = page;
    }
    this.rest.getStaffPoints(page, 10).subscribe(
      (staffPoints) => {
        this.staffPoints=staffPoints;
        this.getPaginationSP();
      }
    );
  }

  getPaginationSP() {
    this.rest.getStaffPointsCount().subscribe(
      res => {
        let spCount = parseInt(res.toString());
        if (spCount > this.spSize) {
          this.spPages = [];
          for (var i = 0; i < Math.ceil(spCount / this.spSize); i++) {
            this.spPages.push(i + 1)
          }
          this.paginationSP = true;
        } else {
          this.spPages = [];
          this.paginationSP = false;
        }
      }
    )
  }

  addStaffPoint() {
    if (this.selectedPoint && this.selectedStaff) {
      if ((this.staffPoints.find(x => x.StaffID == this.selectedStaff.ID)) && (this.staffPoints.find(x => x.PointID == this.selectedPoint.ID))) {
        this.alerts.pushAlert("danger", "ALERTS.STAFFPOINTUSERERR");
      } else {
        let postStaffPoint: StaffPoint = {PointID: this.selectedPoint.ID, StaffID: this.selectedStaff.ID};
        this.rest.postStaffPoint(postStaffPoint).subscribe(
          (postStaffPoint) => {
            this.getStaffPoints();
          }
        );
      }
    } else {
      this.alerts.pushAlert("danger", "ALERTS.STAFFPOINTERR");
    }
  }

  onSelectSP(staffPoint: StaffPoint): void {
    this.selectedSP = staffPoint;
  }

  deletePointStaff() {
    this.rest.deleteStaffPoint(this.selectedSP).subscribe(
      (selectedSP) => {
        this.getStaffPoints();
      }
    );
  }

}
