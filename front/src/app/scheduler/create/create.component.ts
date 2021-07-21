import { Component, OnInit } from '@angular/core';
import { SchedulerService } from '../scheduler.service';
import { BsModalRef } from "ngx-bootstrap/modal";
import { Scheduler } from "../scheduler";
import { BsLocaleService } from 'ngx-bootstrap/datepicker';
import { defineLocale, ptBrLocale } from 'ngx-bootstrap/chronos';

@Component({
  selector: 'app-create',
  templateUrl: './create.component.html',
  styleUrls: ['./create.component.css']
})
export class CreateComponent implements OnInit {
  scheduler: Scheduler;

  constructor(
    public schedulerService: SchedulerService,
    public modalRef: BsModalRef,
    private localeService: BsLocaleService) {
      defineLocale('pt-br', ptBrLocale);    
      this.localeService.use('pt-br');
      
      
      this.scheduler = {};  
    }

  ngOnInit(): void {
  }

  create(){

    this.schedulerService.create(this.scheduler).subscribe(res => {
         console.log('Scheduler created successfully!');
         this.modalRef.hide();
    })
  }

  close() {
    this.modalRef.hide();
  }



}
