import { Component } from '@angular/core';
import { TranslateService } from '@ngx-translate/core';
import { BsDatepickerConfig, BsLocaleService } from 'ngx-bootstrap/datepicker';

import { defineLocale } from 'ngx-bootstrap/chronos';
import { ruLocale } from 'ngx-bootstrap/locale';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})

export class AppComponent {
  title = 'ForteBuisness';
  constructor(
      public translate: TranslateService, 
      public localeService: BsLocaleService) {
    translate.addLangs(["ru", "en"]);
    translate.setDefaultLang('ru');
    defineLocale('ru', ruLocale); 
    let browserLang = translate.getBrowserLang();
    translate.use(browserLang.match(/ru|en/) ? browserLang : 'ru');
    localeService.use(browserLang.match(/ru|en/) ? browserLang : 'ru');
  }
}