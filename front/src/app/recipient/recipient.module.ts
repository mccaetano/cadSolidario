import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { RecipientRoutingModule } from './recipient-routing.module';
import { ViewComponent } from './view/view.component';
import { CreateComponent } from './create/create.component';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { BsDatepickerModule } from 'ngx-bootstrap/datepicker';
import { allIcons, NgxBootstrapIconsModule } from 'ngx-bootstrap-icons';


@NgModule({
  declarations: [
    ViewComponent,
    CreateComponent
  ],
  imports: [
    CommonModule,
    RecipientRoutingModule,
    FormsModule,
    ReactiveFormsModule,
    BsDatepickerModule.forRoot(),
    NgxBootstrapIconsModule.pick(allIcons)
  ]
})
export class RecipientModule { }
