import { Component, OnInit } from '@angular/core';
import { SchedulerService } from '../scheduler.service';
import { BsModalRef } from "ngx-bootstrap/modal";
import { Scheduler } from "../scheduler";

@Component({
  selector: 'app-create',
  templateUrl: './create.component.html',
  styleUrls: ['./create.component.css']
})
export class CreateComponent implements OnInit {
  scheduler: Scheduler;

  constructor(
    public schedulerService: SchedulerService,
    public modalRef: BsModalRef)
    { this.scheduler = {};  }

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
