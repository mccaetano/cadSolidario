import { ChangeDetectorRef, Component, OnInit, TemplateRef } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { Scheduler } from '../scheduler';
import { SchedulerService } from '../scheduler.service';
import { BsModalService, BsModalRef } from "ngx-bootstrap/modal";
import { CreateComponent } from '../create/create.component';
import { UpdateComponent } from '../update/update.component';
import { combineLatest, Subscription } from 'rxjs';


@Component({
  selector: 'app-index',
  templateUrl: './index.component.html',
  styleUrls: ['./index.component.css']
})
export class IndexComponent implements OnInit {
  eventDateSearch?='';
  nameSearch='';
  currentPage=1;
  schedulers: Scheduler[] = [];
  eventDates: Scheduler[] = [];
  modalRef?: BsModalRef
  subscriptions: Subscription[] = [];
  messages: string[] = [];

  constructor(
    public schedulerService: SchedulerService,
    private route: ActivatedRoute,
    public modalService: BsModalService,
    private changeDetectorRef: ChangeDetectorRef) {  }

  ngOnInit(): void {
    this.getEventDates();
  }
  getEventDates(): void {
    this.schedulerService.getEventDates().subscribe((data: Scheduler[]) => {
      this.eventDates = data;
      if (data.length > 0) {
        this.eventDateSearch = data[0]?.eventDate?.toString() ;
      }
      this.getByFilter();
    });
  }

  getByFilter(): void {
    console.log("eventDateSearch: " + this.eventDateSearch);
    var eventDate: Date = new Date('1900-01-01');
    if (this.eventDateSearch != '' && this.eventDateSearch != null) {
      eventDate = new Date(this.eventDateSearch);
    }
    if (this.currentPage <= 0) {
      this.currentPage = 1;
    }
    this.schedulerService.getAll(eventDate, this.nameSearch, this.currentPage).subscribe(
      (data: Scheduler[]) => {
        this.schedulers = data;
        console.log("load data is sucesss");
      }
    );
  }

  changePage(pg: string): void {
    if (pg == "m") {
      if (this.currentPage > 1) {
        this.currentPage -= this.currentPage;
      } else {
        this.currentPage = 1;
      }
    } else {
      if (pg == "p") {
        if (this.schedulers.length > 0) {
          this.currentPage += this.currentPage;
        }
      }
    }
    this.getByFilter();
  }

  openModalCreate() {
    const initialState = {
    };
    this.modalRef = this.modalService.show(CreateComponent, initialState);
    const _combine = combineLatest(
      this.modalRef.onHide,
      this.modalRef.onHidden
    ).subscribe(() => this.changeDetectorRef.markForCheck());

    this.subscriptions.push(
      this.modalRef.onHide.subscribe(() => {
        })
    );
    this.subscriptions.push(
      this.modalRef.onHidden.subscribe(() => {
        this.getEventDates();
        this.unsubscribe();
      })
    );

    this.subscriptions.push(_combine);
  }

  openModalUpdate(id?: number) {
    const initialState = {
      id: id
    };
    console.log("input: " + JSON.stringify( initialState));
    this.modalRef = this.modalService.show(UpdateComponent, {initialState});
    this.modalRef.onHide.pipe()
    const _combine = combineLatest(
      this.modalRef.onHide,
      this.modalRef.onHidden
    ).subscribe(() => this.changeDetectorRef.markForCheck());

    this.subscriptions.push(
      this.modalRef.onHide.subscribe(() => {
        })
    );
    this.subscriptions.push(
      this.modalRef.onHidden.subscribe(() => {
        this.getByFilter();
        this.unsubscribe();
      })
    );

    this.subscriptions.push(_combine);

  }

  unsubscribe() {
    this.subscriptions.forEach((subscription: Subscription) => {
      subscription.unsubscribe();
    });
    this.subscriptions = [];
  }
}
