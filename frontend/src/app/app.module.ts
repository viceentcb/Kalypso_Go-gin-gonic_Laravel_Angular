import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { CarouselModule } from 'ngx-owl-carousel-o';
import { NgxPaginationModule } from 'ngx-pagination';
import { ToastrModule } from 'ngx-toastr';
import { AppComponent } from './app.component';
import { AuthModule } from './auth/auth.module';
import { HomeModule } from './home/home.module';
import { ProductModule } from './product/products.module';
import {
  FooterComponent,
  HeaderComponent,
  SharedModule
} from './shared';
// import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { AppRoutingModule } from './app-routing.module';
import { CoreModule } from './core/core.module';
import { HttpClientModule } from '@angular/common/http';
//import { from } from 'rxjs';

@NgModule({
  declarations: [AppComponent, FooterComponent, HeaderComponent],
  imports: [
    BrowserModule,
    CoreModule,
    SharedModule,
    HomeModule,
    AuthModule,
    AppRoutingModule,
    ProductModule,
    CarouselModule,
    NgxPaginationModule,
    // BrowserAnimationsModule,
    HttpClientModule,

    ToastrModule.forRoot()
  ],
  providers: [],
  bootstrap: [
    AppComponent,
  ]
})
export class AppModule {}
