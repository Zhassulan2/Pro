import { Component, OnInit } from '@angular/core';
import { OAuthService } from "angular-oauth2-oidc";

@Component({
  selector: 'app-dictionaries',
  templateUrl: './dictionaries.component.html',
  styleUrls: ['./dictionaries.component.css']
})
export class DictionariesComponent implements OnInit {

  constructor(
    private oauthService: OAuthService, 
  ) { }

  ngOnInit() {
  }
}
