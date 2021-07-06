package models

import (
	"log"
)

type Contact struct {
	Phone    string `json:"phone,omitempty"`
	CelPhone string `json:"celPhone,omitempty"`
}

type Document struct {
	Rg   string `json:"rg,omitempty"`
	Cpf  string `json:"cpf,omitempty"`
	Cpts string `json:"cpts,omitempty"`
	Pis  string `json:"pis,omitempty"`
}

type Dependent struct {
	Name     string `json:"name,omitempty"`
	Document string `json:"document,omitempty"`
}

type Recipient struct {
	Id          int64       `json:"id,omitempty"`
	Name        string      `json:"name,omitempty"`
	Birthdate   string      `json:"birthDate,omitempty"`
	Address     string      `json:"address,omitempty"`
	Work        string      `json:"work,omitempty"`
	Documents   Document    `json:"documents,omitempty"`
	Contacts    Contact     `json:"contacts,omitempty"`
	Dependents  []Dependent `json:"dependents,omitempty"`
	Retiree     bool        `json:"retiree,omitempty"`
	RentPay     bool        `json:"rentPay,omitempty"`
	Working     int32       `json:"working,omitempty"`
	HomePeaples int32       `json:"homePeaples,omitempty"`
	Milks       int32       `json:"milks,omitempty"`
	Babys       int32       `json:"babys,omitempty"`
	Boys        int32       `json:"boys,omitempty"`
	Girls       int32       `json:"girls,omitempty"`
	HelpFamily  bool        `json:"helpFamily,omitempty"`
	Active      bool        `json:"Active,omitempty"`
}

func RecipientGetById(id int64) (Recipient, error) {

	row := DB.QueryRow(`select
			id,
			nome,
			to_char(data_nacimento, 'YYYY-MM-DD') data_nacimento,
			profissao,
			documento_rg,
			documento_cpf,
			documento_cpts,
			documento_pis,
			contato_tel,
			contato_cel,
			endereco,
			dependente_nome,
			dependente_documento,
			dependente_1_nome,
			dependente_2_nome,
			paga_aluguel,
			aposentado,
			bolsa_familia,
			qtde_pessoas_trabalham,
			qtde_pessoas_casa,
			qtde_leite,
			qtde_bebe,
			qtde_meninos,
			qtde_meninas,
			ativo
		from
			public.tbbeneficiario 
		where 
			id  = $1`,
		id)

	var recipient Recipient
	name1, name2, name3, document := "", "", "", ""

	row.Scan(&recipient.Id, &recipient.Name, &recipient.Birthdate, &recipient.Work,
		&recipient.Documents.Rg, &recipient.Documents.Cpf, &recipient.Documents.Cpts,
		&recipient.Documents.Pis, &recipient.Contacts.Phone, &recipient.Contacts.CelPhone,
		&recipient.Address, &name1, &document,
		&name2, &name3, &recipient.RentPay,
		&recipient.Retiree, &recipient.HelpFamily, &recipient.Working, &recipient.HomePeaples,
		&recipient.Milks, &recipient.Babys, &recipient.Boys, &recipient.Girls, &recipient.Active)

	recipient.Dependents = []Dependent{}
	recipient.Dependents = append(recipient.Dependents, Dependent{Name: name1, Document: document})
	recipient.Dependents = append(recipient.Dependents, Dependent{Name: name2})
	recipient.Dependents = append(recipient.Dependents, Dependent{Name: name3})

	return recipient, nil
}

func RecipientGetByFilter(name string, limit int32, skip int32) ([]Recipient, error) {
	offset := limit * (skip - 1)
	log.Println("Repo: (RecipientGetByFilter) Get - name: ", name)
	rows, err := DB.Query(`select
				id,
				nome,
				to_char(data_nacimento, 'YYYY-MI-DD') data_nascimento,
				profissao,
				documento_rg,
				documento_cpf,
				documento_cpts,
				documento_pis,
				contato_tel,
				contato_cel,
				endereco,
				dependente_nome,
				dependente_documento,
				dependente_1_nome,
				dependente_2_nome,
				paga_aluguel,
				aposentado,
				bolsa_familia,
				qtde_pessoas_trabalham,
				qtde_pessoas_casa,
				qtde_leite,
				qtde_bebe,
				qtde_meninos,
				qtde_meninas,
				ativo
			from
				public.tbbeneficiario 
			where 
				nome  like '%' || $1 || '%'
			limit $2 offset $3`,
		name,
		limit,
		offset)
	if err != nil {
		log.Println("Repo: (RecipientGetByFilter) Get - Erro:", err)
		return nil, err
	}
	defer rows.Close()

	var recipients []Recipient = []Recipient{}
	for rows.Next() {
		var recipient Recipient
		name1, name2, name3, document := "", "", "", ""
		recipient.Dependents = []Dependent{}
		rows.Scan(&recipient.Id, &recipient.Name, &recipient.Birthdate, &recipient.Work,
			&recipient.Documents.Rg, &recipient.Documents.Cpf, &recipient.Documents.Cpts,
			&recipient.Documents.Pis, &recipient.Contacts.Phone, &recipient.Contacts.CelPhone,
			&recipient.Address, &name1, &document,
			&name2, &name3, &recipient.RentPay,
			&recipient.Retiree, &recipient.HelpFamily, &recipient.Working, &recipient.HomePeaples,
			&recipient.Milks, &recipient.Babys, &recipient.Boys, &recipient.Girls, &recipient.Active)

		recipient.Dependents = []Dependent{}
		recipient.Dependents = append(recipient.Dependents, Dependent{Name: name1, Document: document})
		recipient.Dependents = append(recipient.Dependents, Dependent{Name: name2})
		recipient.Dependents = append(recipient.Dependents, Dependent{Name: name3})
		recipients = append(recipients, recipient)
	}

	return recipients, nil
}

func RecipientPost(recipient Recipient) (int64, error) {
	var id int64

	DB.QueryRow(`INSERT INTO public.tbbeneficiario
		(nome, data_nacimento, profissao, documento_rg, documento_cpf, documento_cpts, 
		documento_pis, contato_tel, contato_cel, endereco, dependente_nome, dependente_documento, dependente_1_nome, dependente_2_nome, 
		paga_aluguel, aposentado, bolsa_familia, qtde_pessoas_trabalham, qtde_pessoas_casa, qtde_leite, qtde_bebe, qtde_meninos,
		qtde_meninas, ativo)
		VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, true);
		RETURNING id`,
		recipient.Name,
		recipient.Birthdate,
		recipient.Work,
		recipient.Documents.Rg,
		recipient.Documents.Cpf,
		recipient.Documents.Cpts,
		recipient.Documents.Pis,
		recipient.Contacts.Phone,
		recipient.Contacts.CelPhone,
		recipient.Address,
		recipient.Dependents[0].Name,
		recipient.Dependents[0].Document,
		recipient.Dependents[1].Name,
		recipient.Dependents[2].Name,
		recipient.RentPay,
		recipient.Retiree,
		recipient.HelpFamily,
		recipient.Working,
		recipient.HomePeaples,
		recipient.Milks,
		recipient.Babys,
		recipient.Boys,
		recipient.Girls).Scan(&id)

	return id, nil
}

func RecipientPut(id int64, recipient Recipient) (Recipient, error) {

	_, err := DB.Exec(`UPDATE public.tbbeneficiario
		SET nome=$1, data_nacimento=$2, profissao=$3, documento_rg=$4, documento_cpf=$5, documento_cpts=$6, documento_pis=$7, 
		contato_tel=$8, contato_cel=$9, endereco=$10, dependente_nome=$11, dependente_documento=$12, dependente_1_nome=$13, 
		dependente_2_nome=$14, paga_aluguel=$15, aposentado=$16, bolsa_familia=$17, qtde_pessoas_trabalham=$18, 
		qtde_pessoas_casa=$19, qtde_leite=$20, qtde_bebe=$21, qtde_meninos=$22, qtde_meninas=$23, ativo=$24
		WHERE id=$1`,
		id,
		recipient.Name,
		recipient.Birthdate,
		recipient.Work,
		recipient.Documents.Rg,
		recipient.Documents.Cpf,
		recipient.Documents.Cpts,
		recipient.Documents.Pis,
		recipient.Contacts.Phone,
		recipient.Contacts.CelPhone,
		recipient.Address,
		recipient.Dependents[0].Name,
		recipient.Dependents[0].Document,
		recipient.Dependents[1].Name,
		recipient.Dependents[2].Name,
		recipient.RentPay,
		recipient.Retiree,
		recipient.HelpFamily,
		recipient.Working,
		recipient.HomePeaples,
		recipient.Milks,
		recipient.Babys,
		recipient.Boys,
		recipient.Girls)
	if err != nil {
		log.Println("Erro lendo datados:", err)
		return Recipient{}, err
	}

	retorno, err := RecipientGetById(id)
	if err != nil {
		log.Println("Erro lendo datados:", err)
		return Recipient{}, err
	}
	return retorno, err
}
