create table system_intake (
    id uuid PRIMARY KEY not null,
    eua_user_id text not null,
    requester text,
    component text,
    business_owner text,
    business_owner_component text,
    product_manager text,
    product_manager_component text,
    has_isso bool,
    isso text,
    trb_collaborator text,
    oit_security_collaborator text,
    ea_collaborator text,
    project_name text,
    existing_funding bool,
    funding_source text,
    business_need text,
    solution text,
    process_status text,
    ea_support_request bool,
    existing_contract text
);
