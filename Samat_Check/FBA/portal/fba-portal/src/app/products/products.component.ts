import { Component, OnInit } from '@angular/core';
import { Product } from "../models/product";
import { Category } from "../models/category";
import { RestService } from '../rest.service';

@Component({
  selector: 'app-products',
  templateUrl: './products.component.html',
  styleUrls: ['./products.component.css']
})
export class ProductsComponent implements OnInit {
    //New Product
    newName: string;
    newBarcode: string;
    newCategory: string = "";

    categories: Array<Category> = [];
    products: Array<Product> = [];
    selectedProduct: Product;
    clientSecret: string;
    productsScope: boolean = false;
    searchScope: boolean = false;
    paginationProducts: boolean = false;
    productPages = [];
    productSize: number = 10;
    currentPage: number = 1;
  
    constructor(
      private rest: RestService
    ) { }
  
    ngOnInit() {
      this.getProducts();
      this.getCategories();
    }

    getCategories() {
      this.categories = [];
      this.rest.getCategoriesCount().subscribe(
        res => {
          let categoriesCount = parseInt(res.toString());
          this.rest.getCategories(1, categoriesCount).subscribe(
            (categories) => {
              this.categories=categories;
            }
          );
        }
      )
    }

    refreshProduct() {
      this.getProducts();
    }
  
    onSelect(product: Product): void {
      this.selectedProduct = product;
    }
  
    getPagination() {
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
  
    getProducts(page?: number, name?: string, barcode?: string) {
      this.products = [];
      if (page) {
        this.currentPage = page;
      }
      this.rest.getProducts(page, 10, name, barcode).subscribe(
        (products) => {
          this.products=products;
          this.productsScope = true;
          this.getPagination();
        }
      );
    }
  
    createProduct() {
      let postProduct: Product = {Name: this.newName, Barcode: this.newBarcode, CategoryID: this.newCategory};
      this.rest.postProducts(postProduct).subscribe(
        (postProduct) => {
          this.getProducts();
          this.newName = null;
          this.newBarcode = null;
          this.newCategory = "";
        }
      );
    }
  
    updateProduct() {
      this.rest.putProducts(this.selectedProduct).subscribe(
        (selectedProduct) => {
          this.getProducts();
        }
      );
    }
  
    deleteProduct() {
      if (confirm("Вы уверены что хотите удалить точку " + this.selectedProduct.Name + "?")) {
        this.rest.deleteProducts(this.selectedProduct).subscribe(
          (selectedProduct) => {
            this.getProducts();
          }
        );
      }
    }
  }