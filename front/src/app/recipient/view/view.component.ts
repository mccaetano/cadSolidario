import { Component, OnInit, TemplateRef } from '@angular/core';
import { FormBuilder, FormGroup } from '@angular/forms';
import { Router } from '@angular/router';
import { throwError } from 'rxjs';
import { Recipient } from '../recipient';
import { RecipientService } from '../recipient.service';
import { BsModalService, BsModalRef } from 'ngx-bootstrap/modal';
import { BsLocaleService } from 'ngx-bootstrap/datepicker';
import { defineLocale, ptBrLocale } from 'ngx-bootstrap/chronos';

@Component({
  selector: 'app-view',
  templateUrl: './view.component.html',
  styleUrls: ['./view.component.css']
})
export class ViewComponent implements OnInit {

  nameSearch: string = '';
  recipients: Recipient[] = [];
  form: FormGroup;
  selRecipientId: number = 0;
  modalRef: BsModalRef | any;

  constructor(public recipientService: RecipientService, 
    private formBuilder: FormBuilder, 
    private router: Router, 
    private modalService: BsModalService,
    private localeService: BsLocaleService) {
    
    defineLocale('pt-br', ptBrLocale);
    this.localeService.use('pt-br');

    this.form = this.formBuilder.group({
      id: [''],
      name: [''],
      birthDate: [new Date()],
      work: [''],
      phone: [''],
      celPhone: [''],
      address: [''],
      documentRg: [''],
      documentCpf: [''],
      documentCpts: [''],
      documentPis: [''],
      dependents0Name: [''],
      dependents0Doc: [''],
      dependents1Name: [''],
      dependents2Name: [''],
      retiree: [''],
      rentPay: [''],
      working: [''],
      homePeaples: [''],
      milks: [''],
      babys: [''],
      boys: [''],
      girls: [''],
      helpFamily: [''],
      active: ['']
    });
  }

  ngOnInit(): void {
    this.getByFilter();
  }

  getByFilter() {
    console.log("name:" + this.nameSearch);
    this.recipientService.getByFilter(this.nameSearch).subscribe(
      (data: Recipient[]) => {
        this.recipients = data;
      }
    );
  }

  getById(id: number) {
    this.recipientService.getById(id).subscribe(
      (data: Recipient) => {
        var value: any = {
          id: data.id ?? 0,
          name: data.name ?? "",
          birthDate: new Date(data?.birthDate ?? '1900-01-01'),
          work: data.work ?? "",
          phone: data.contacts?.phone ?? "",
          celPhone: data.contacts?.celPhone ?? "",
          address: data.address ?? "",
          documentRg: data.documents?.rg ?? "",
          documentCpf: data.documents?.cpf ?? "",
          documentCpts: data.documents?.cpts ?? "",
          documentPis: data.documents?.pis ?? "",
          dependents0Name: data.dependents?.[0].name ?? "",
          dependents0Doc: data.dependents?.[0].document ?? "",
          dependents1Name: data.dependents?.[1].name ?? "",
          dependents2Name: data.dependents?.[2].name ?? "",
          retiree: data.retiree ?? false,
          rentPay: data.rentPay ?? false,
          working: data.working ?? 0,
          homePeaples: data.homePeaples ?? 0,
          milks: data.milks ?? 0,
          babys: data.babys ?? 0,
          boys: data.boys ?? 0,
          girls: data.girls ?? 0,
          helpFamily: data.helpFamily ?? false,
          active: data.active ?? true
        };
        console.log("Recipient View: " + JSON.stringify(value));
        this.form.setValue(value);
      }
    );
    this.selRecipientId = id;
  }

  create() {
    this.router.navigate(['/recipient/create']);
  }

  openModal(template: TemplateRef<any>) {
    this.modalRef = this.modalService.show(template, {class: 'modal-sm'});
  }
 
  confirm(): void {
    this.modalRef.hide();
    var recipient: Recipient = {
      name: this.form.get("name")?.value ?? "",
      birthDate: new Date(this.form.get("birthDate")?.value == "" ? "1900-01-01": this.form.get("birthDate")?.value).toISOString().substring(0,10),
      work: this.form.get("work")?.value ?? "",
      address: this.form.get("address")?.value ?? "",
      contacts: {
        phone: this.form.get("phone")?.value ?? "",
        celPhone: this.form.get("celPhone")?.value ?? ""
      },
      documents: {
        rg: this.form.get("documentRg")?.value ?? "",
        cpf: this.form.get("documentCpf")?.value ?? "",
        cpts: this.form.get("documentCpts")?.value ?? "",
        pis: this.form.get("documentPis")?.value ?? ""
      },
      dependents: [
        {
          name: this.form.get("dependents0Name")?.value ?? "",
          document: this.form.get("dependents0Doc")?.value ?? "",
        },
        {
          name: this.form.get("dependents1Name")?.value ?? ""
        },
        {
          name: this.form.get("dependents02Name")?.value ?? ""
        },
      ],
      retiree: this.form.get("retiree")?.value ?? false,
      rentPay: this.form.get("rentPay")?.value ?? false,
      working: Number(this.form.get("working")?.value ?? 0),
      active: this.form.get("active")?.value ?? true,
      milks: Number(this.form.get("milks")?.value ?? 0),
      babys: Number(this.form.get("babys")?.value ?? 0),
      boys: Number(this.form.get("boys")?.value ?? 0),
      girls: Number(this.form.get("girls")?.value ?? 0),
      helpFamily: this.form.get("helpFamily")?.value ?? false,
      homePeaples: Number(this.form.get("homePeaples")?.value ?? 0)
    };
    console.log("Update: id: " + this.selRecipientId + ", body" + JSON.stringify(recipient));
    this.recipientService.update(this.selRecipientId, recipient).subscribe( res => {
      console.log("Finins update");
      this.selRecipientId = 0;    
      this.form.reset;
    });    

  }
 
  decline(): void {
    this.modalRef.hide();
  }
  
  errorHandler(error: any) {
    let errorMessage = '';
    if(error.error instanceof ErrorEvent) {
      errorMessage = error.error.message;
    } else {
      errorMessage = `Error Code: ${error.status}\nMessage: ${error.message}`;
    }
    console.error(errorMessage);    
    return throwError(errorMessage);
 }
}
