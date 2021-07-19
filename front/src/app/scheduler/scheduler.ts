import { StatusList } from "./statusList";

export interface Document {
    rg?: string;
    cpf?: string;
    cpts?: string;
    pis?: string;
}


export interface Contact {
    phone?: string;
    celPhone?: string;
}

export interface Scheduler {
    id?: number;
    eventDate?: Date;
    name?: string;
    documents?: Document;
    contacts?: Contact;
    milks?: number;
    status?: StatusList;
}
