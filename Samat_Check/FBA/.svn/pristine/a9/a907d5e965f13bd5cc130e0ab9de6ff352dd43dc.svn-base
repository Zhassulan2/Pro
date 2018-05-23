import { Component, OnInit } from '@angular/core';
import { RestService } from '../rest.service';
import { Category } from "../models/category";
import { PermissionCheckService } from '../permission-check.service';

@Component({
  selector: 'app-category',
  templateUrl: './category.component.html',
  styleUrls: ['./category.component.css']
})
export class CategoryComponent implements OnInit {
  //New Category
  newName: string;
  newShort: string;
  
  testname: string;
  categories: Array<Category> = [];
  selectedCategory: Category;
  clientSecret: string;
  categoriesScope: boolean = false;
  searchScope: boolean = false;
  paginationCategories: boolean = false;
  categoryPages = [];
  categorySize: number = 10;
  currentPage: number = 1;
  
  constructor(
    private rest: RestService,
    private check: PermissionCheckService
  ) { }

  ngOnInit() {
    this.getCategories();
  }

  refreshCategory() {
    this.getCategories();
  }

  onSelect(category: Category): void {
    this.selectedCategory = category;
  }

  getPagination() {
    this.rest.getCategoriesCount().subscribe(
      res => {
        let categoryCount = parseInt(res.toString());
        if (categoryCount > this.categorySize) {
          this.categoryPages = [];
          for (var i = 0; i < Math.ceil(categoryCount / this.categorySize); i++) {
            this.categoryPages.push(i + 1)
          }
          this.paginationCategories = true;
        } else {
          this.categoryPages = [];
          this.paginationCategories = false;
        }
      }
    )
  }

  getCategories(page?: number) {
    this.categories = [];
    if (page) {
      this.currentPage = page;
    }
    this.rest.getCategories(page, 10).subscribe(
      (categories) => {
        this.categories=categories;
        this.categoriesScope = true;
        this.getPagination();
      }
    );
  }

  createCategory() {
    let postCategory: Category = {Name: this.newName, ShortName: this.newShort};
    this.rest.postCategories(postCategory).subscribe(
      (postCategory) => {
        this.getCategories();
        this.newName = null;
      }
    );
  }

  updateCategory() {
    this.rest.putCategories(this.selectedCategory).subscribe(
      (selectedCategory) => {
        this.getCategories();
      }
    );
  }

  deleteCategory() {
    if (confirm("Вы уверены что хотите удалить точку " + this.selectedCategory.Name + "?")) {
      this.rest.deleteCategories(this.selectedCategory).subscribe(
        (selectedCategory) => {
          this.getCategories();
        }
      );
    }
  }
}

