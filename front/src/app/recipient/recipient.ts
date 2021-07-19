
export interface Document {
	rg?:   string;
	cpf?:  string;
	cpts?: string;
	pis?:  string;
}

export interface Contact {
	phone?:    string;
	celPhone?: string;
}

export interface Dependent {
	name?:     string;
	document?: string;
}

export interface Recipient {
  id?: number;
  name?: string;
  birthDate?: string;
  address?: string;
  work?: string;
  documents?: Document;
  contacts?: Contact;
  dependents?: Dependent[];
  retiree?: boolean;
  rentPay?: boolean;
  working?: number;
  homePeaples?: number;
  milks?: number;
  babys?: number;
  boys?: number;
  girls?: number;
  helpFamily?: boolean;
  active?: boolean;
}
