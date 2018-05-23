import { Component, OnInit } from '@angular/core';
import { OAuthService } from "angular-oauth2-oidc";
import { RestService } from '../rest.service';
import { Staff } from "../models/staff";
import { Role } from "../models/role";

@Component({
  selector: 'app-staff',
  templateUrl: './staff.component.html',
  styleUrls: ['./staff.component.css']
})
export class StaffComponent implements OnInit {
  //New Staff
  newName: string;
  newRole: string = "";

  roles: Array<Role> = [];
  staffs: Array<Staff> = [];
  selectedStaff: Staff;
  clientSecret: string;
  staffsScope: boolean = false;
  searchScope: boolean = false;
  paginationStaffs: boolean = false;
  staffPages = [];
  staffSize: number = 10;
  currentPage: number = 1;
  searchName: string;

  constructor(
    private oauthService: OAuthService,
    private rest: RestService
  ) { }

  ngOnInit() {
    this.getStaffs();
    this.getRoles();
  }

  refreshStaff() {
    this.searchName = null;
    this.getStaffs();
  }

  onSelect(staff: Staff): void {
    this.selectedStaff = staff;
  }

  getRoles() {
    this.roles = [];
    this.rest.getRoles().subscribe(
      (roles) => {
        this.roles=roles;
      }
    );
  }

  getPagination() {
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

  getStaffs(page?: number) {
    this.staffs = [];
    if (page) {
      this.currentPage = page;
    }
    this.rest.getStaffs(page, 10).subscribe(
      (staffs) => {
        this.staffs=staffs;
        this.staffsScope = true;
        this.getPagination();
      }
    );
  }

  createStaff() {
    let postStaff: Staff = {Name: this.newName, RoleID: this.newRole, ParentID: this.userId};
    this.rest.postStaffs(postStaff).subscribe(
      (postStaff) => {
        this.getStaffs();
        this.newName = null;
        this.newRole = "";
      }
    );
  }

  updateStaff() {
    this.rest.putStaffs(this.selectedStaff).subscribe(
      (selectedStaff) => {
        this.getStaffs();
      }
    );
  }

  deleteStaff() {
    if (confirm("Вы уверены что хотите удалить точку " + this.selectedStaff.Name + "?")) {
      this.rest.deleteStaffs(this.selectedStaff).subscribe(
        (selectedStaff) => {
          this.getStaffs();
        }
      );
    }
  }

  get userId() {
    var claims = this.oauthService.getIdentityClaims();
    if (!claims) return null;
    return claims['ID'];
  }
}