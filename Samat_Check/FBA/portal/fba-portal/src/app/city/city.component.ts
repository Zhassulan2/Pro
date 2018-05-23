import { Component, OnInit } from '@angular/core';
import { RestService } from '../rest.service';
import { City } from "../models/city";

@Component({
  selector: 'app-city',
  templateUrl: './city.component.html',
  styleUrls: ['./city.component.css']
})
export class CityComponent implements OnInit {
  cities: Array<City> = [];
  selectedCity: City;
  clientSecret: string;
  citiesScope: boolean = false;
  searchScope: boolean = false;
  paginationCities: boolean = false;
  cityPages = [];
  citySize: number = 10;
  currentPage: number = 1;
  
  constructor(
    private rest: RestService
  ) { }

  ngOnInit() {
    this.getCities();
  }

  refreshCity() {
    this.getCities();
  }

  onSelect(city: City): void {
    this.selectedCity = city;
  }

  getPagination() {
    this.rest.getCitiesCount().subscribe(
      res => {
        let cityCount = parseInt(res.toString());
        if (cityCount > this.citySize) {
          this.cityPages = [];
          for (var i = 0; i < Math.ceil(cityCount / this.citySize); i++) {
            this.cityPages.push(i + 1)
          }
          this.paginationCities = true;
        } else {
          this.cityPages = [];
          this.paginationCities = false;
        }
      }
    )
  }

  getCities(page?: number) {
    this.cities = [];
    if (page) {
      this.currentPage = page;
    }
    this.rest.getCities(page, 10).subscribe(
      (cities) => {
        this.cities=cities;
        this.citiesScope = true;
        this.getPagination();
      }
    );
  }

  createCity(name: string) {
    let postCity: City = {Name: name};
    this.rest.postCities(postCity).subscribe(
      (postCity) => {
        this.getCities();
      }
    );
  }

  updateCity() {
    this.rest.putCities(this.selectedCity).subscribe(
      (selectedCity) => {
        this.getCities();
      }
    );
  }

  deleteCity() {
    if (confirm("Вы уверены что хотите удалить точку " + this.selectedCity.Name + "?")) {
      this.rest.deleteCities(this.selectedCity).subscribe(
        (selectedCity) => {
          this.getCities();
        }
      );
    }
  }
}
