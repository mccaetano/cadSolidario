import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup } from '@angular/forms';
import { Router } from '@angular/router';
import { throwError } from 'rxjs';
import { catchError } from 'rxjs/operators';
import { Recipient } from '../recipient';
import { RecipientService } from '../recipient.service';

@Component({
  selector: 'app-create',
  templateUrl: './create.component.html',
  styleUrls: ['./create.component.css']
})
export class CreateComponent implements OnInit {

  form: FormGroup;
  constructor(
    public formBuilder: FormBuilder,
    private router: Router,
    private recipientService: RecipientService) {
    this.form = this.formBuilder.group({
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
  }

  criar() {
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
    console.log("Create: " + JSON.stringify(recipient));
    this.recipientService.create(recipient).pipe(
      catchError(this.errorHandler)
    );
    this.form.reset();
    this.router.navigate(['/recipient']);
  }


  errorHandler(error: any) {
    let errorMessage = '';
    if(error.error instanceof ErrorEvent) {
      errorMessage = error.error.message;
    } else {
      errorMessage = `Error Code: ${error.status}\nMessage: ${error.message}`;
    }
    return throwError(errorMessage);
 }
}
