import { Component, OnInit } from '@angular/core';
import { RestService } from '../rest.service';
import { Company } from "../models/company";
import { PermissionCheckService } from '../permission-check.service';

@Component({
  selector: 'app-company',
  templateUrl: './company.component.html',
  styleUrls: ['./company.component.css']
})
export class CompanyComponent implements OnInit {
  //New Company
  newName:string;
  newAddress:string;
  newBIN:string;
  newContact:string;

  companies: Array<Company> = [];
  selectedCompany: Company;
  clientSecret: string;
  companiesScope: boolean = false;
  searchScope: boolean = false;
  paginationCompanies: boolean = false;
  companyPages = [];
  companySize: number = 10;
  currentPage: number = 1;
  
  constructor(
    private rest: RestService,
    private check: PermissionCheckService
  ) { }

  ngOnInit() {
    this.getCompanies();
  }

  refreshCompany() {
    this.getCompanies();
  }

  onSelect(company: Company): void {
    this.selectedCompany = company;
  }

  getPagination() {
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

  getCompanies(page?: number) {
    this.companies = [];
    if (page) {
      this.currentPage = page;
    }
    this.rest.getCompanies(page, 10).subscribe(
      (companies) => {
        this.companies=companies;
        this.companiesScope = true;
        this.getPagination();
      }
    );
  }

  createCompany() {
    let postCompany: Company = {Name: this.newName, Bin: this.newBIN, Address: this.newAddress, Contact:this.newContact};
    this.rest.postCompanies(postCompany).subscribe(
      (postCompany) => {
        this.getCompanies();
        this.newName = null;
        this.newAddress = null;
        this.newBIN = null;
        this.newContact = null;
      }
    );
  }

  updateCompany() {
    this.rest.putCompanies(this.selectedCompany).subscribe(
      (selectedCompany) => {
        this.getCompanies();
      }
    );
  }

  deleteCompany() {
    if (confirm("Вы уверены что хотите удалить точку " + this.selectedCompany.Name + "?")) {
      this.rest.deleteCompanies(this.selectedCompany).subscribe(
        (selectedCompany) => {
          this.getCompanies();
        }
      );
    }
  }
}

