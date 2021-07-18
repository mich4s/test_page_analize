import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppComponent } from './app.component';
import {HttpClientModule} from "@angular/common/http";
import { BooleanPipe } from './pipes/boolean.pipe';
import {FormsModule} from "@angular/forms";
import { DebounceClickDirective } from './directives/debounce-click.directive';

@NgModule({
  declarations: [
    AppComponent,
    BooleanPipe,
    DebounceClickDirective
  ],
  imports: [
    BrowserModule,
    HttpClientModule,
    FormsModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
