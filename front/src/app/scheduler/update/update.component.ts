import { Component, Input, OnInit } from '@angular/core';
import { BsModalRef } from 'ngx-bootstrap/modal';
import { Scheduler } from '../scheduler';
import { SchedulerService } from '../scheduler.service';
import { StatusListService } from '../status-list.service';
import { StatusList } from '../statusList';

@Component({
  selector: 'app-update',
  templateUrl: './update.component.html',
  styleUrls: ['./update.component.css']
})
export class UpdateComponent implements OnInit {

  statusCurrent: number = 0;
  statusList: StatusList[] = [];
  @Input() id: number =0;

  constructor(private modalRef: BsModalRef,
    private statusListService: StatusListService,
    private scehdulerService: SchedulerService) {
      this.statusListService.getAll().subscribe(
        (data: StatusList[]) => {
          this.statusList = data;
        }
      );
     }

  ngOnInit(): void {
  }


  close() {
    this.modalRef.hide();
  }

  setStatus() {
    var scheduler: Scheduler = {
      status: {
        id: Number(this.statusCurrent)
      }
    };
    console.log('setStatus: '+ JSON.stringify(scheduler))
    this.scehdulerService.update(this.id, scheduler).subscribe();
    this.close();
  }
}
