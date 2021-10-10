<script>
    import { Row, Col, Container } from "sveltestrap";
    import { createEventDispatcher } from "svelte";
    import { createForm } from "svelte-forms-lib";
    import * as yup from "yup";

    export let sData = "";
    export let token = "";
    export let adminrule_id = "";
    export let adminrule_name = "";
    export let adminrule_rule_field = "";
    export let adminrule_create_field = "";
    export let adminrule_update_field = "";

    let css_loader = "display: none;";
    let msgloader = "";

    let dispatch = createEventDispatcher();
    const schema = yup.object().shape({
        admin_name_field: yup
            .string()
            .required("Name is Required")
            .matches(
                /^[a-zA-z0-9 ]+$/,
                "Name must Character A-Z or a-z or 1-9 or space"
            )
            .min(4, "Name must be at least 4 Character")
            .max(70, "Name must be at most 4 Character"),
    });
    const { form, errors, handleChange, handleSubmit } = createForm({
        initialValues: {
            admin_name_field: adminrule_name,
        },
        validationSchema: schema,
        onSubmit: (values) => {
            SaveTransaksi(values.admin_name_field);
        },
    });
    $: {
        if ($errors.admin_name_field) {
            alert($errors.admin_name_field);
            form.admin_name_field = "";
        }
    }

    const BackHalaman = () => {
        dispatch("handleBackHalaman", "call");
    };

    async function SaveTransaksi(name) {
        const res = await fetch("/api/saveadminrule", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                sdata: sData,
                page: "ADMINRULE-SAVE",
                idrule: adminrule_id,
                nama: name,
            }),
        });
        const json = await res.json();

        if (json.status == 200) {
            msgloader = json.message;
        } else if (json.status == 403) {
            alert(json.message);
        } else {
            msgloader = json.message;
        }
        setTimeout(function () {
            css_loader = "display: none;";
        }, 1000);
    }
    async function Updateconfig() {
        let flag = false;
        let msg = "";
        css_loader = "display: inline-block;";
        msgloader = "Sending...";
        if (adminrule_rule_field == "") {
            flag = true;
            msg += "The List is required\n";
        }

        if (flag == false) {
            const res = await fetch("/api/saveadminruleconf", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: sData,
                    page: "ADMINRULE-SAVE",
                    idrule: adminrule_id,
                    rule: adminrule_rule_field.toString(),
                }),
            });
            const json = await res.json();

            if (json.status == 200) {
                msgloader = json.message;
            } else if (json.status == 403) {
                alert(json.message);
            } else {
                msgloader = json.message;
            }
            setTimeout(function () {
                css_loader = "display: none;";
            }, 1000);
        } else {
            alert(msg);
        }
    }
</script>

<div id="loader" style="margin-left:50%;{css_loader}">
    {msgloader}
</div>
<Container fluid style="margin-top: 70px;">
    <Row>
        <Col>
            <button
                on:click={() => {
                    BackHalaman();
                }}
                class="btn btn-dark"
                style="border-radius: 0px;"
            >
                Back
            </button>
        </Col>
    </Row>
    <div class="clearfix" />
    <Row>
        <Col xs="3">
            <div class="card" style="border-radius: 0px;margin-top:10px;">
                <div class="card-header" style="">
                    Admin Rule / {sData}
                    <div class="float-end">
                        <button
                            on:click={() => {
                                handleSubmit();
                            }}
                            class="btn btn-warning"
                            style="border-radius: 0px;"
                        >
                            Save
                        </button>
                    </div>
                </div>
                <div class="card-body" style="height:500px;">
                    <div class="mb-3">
                        <label for="exampleForm" class="form-label">Name</label>
                        <input
                            on:change={handleChange}
                            bind:value={$form.admin_name_field}
                            invalid={$errors.admin_name_field.length > 0}
                            type="text"
                            maxlength="70"
                            class="form-control"
                            placeholder="Name"
                            aria-label="Name"
                        />
                    </div>

                    <div class="mb-3">
                        <table>
                            <tr>
                                <td style="font-size: 11px;">Create</td>
                                <td style="font-size: 11px;">:</td>
                                <td style="font-size: 11px;"
                                    >{adminrule_create_field}</td
                                >
                            </tr>
                            <tr>
                                <td style="font-size: 11px;">Update</td>
                                <td style="font-size: 11px;">:</td>
                                <td style="font-size: 11px;"
                                    >{adminrule_update_field}</td
                                >
                            </tr>
                        </table>
                    </div>
                </div>
            </div>
        </Col>
        <Col xs="9">
            <div class="card" style="border-radius: 0px;margin-top:10px;">
                <div class="card-header" style="">
                    Setting Pasaran Online
                    <div class="float-end">
                        <button
                            on:click={() => {
                                Updateconfig();
                            }}
                            class="btn btn-warning"
                            style="border-radius: 0px;"
                        >
                            Save
                        </button>
                    </div>
                </div>
                <div class="card-body" style="height:550px;">
                    <table class="table">
                        <thead>
                            <tr>
                                <th colspan="2">DASHBOARD</th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr>
                                <td width="1%">
                                    <input
                                        bind:group={adminrule_rule_field}
                                        type="checkbox"
                                        value="DASHBOARD-VIEW"
                                    />
                                </td>
                                <td width="*">VIEW</td>
                            </tr>
                        </tbody>
                    </table>
                    <table class="table">
                        <thead>
                            <tr>
                                <th colspan="2">PERIODE</th>
                                <th colspan="2">PREDIKSI</th>
                                <th colspan="2">REPORT WINLOSE</th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr>
                                <td width="1%">
                                    <input
                                        bind:group={adminrule_rule_field}
                                        type="checkbox"
                                        value="PERIODE-VIEW"
                                    />
                                </td>
                                <td width="*">VIEW</td>
                                <td width="1%">
                                    <input
                                        bind:group={adminrule_rule_field}
                                        type="checkbox"
                                        value="PREDIKSI-VIEW"
                                    />
                                </td>
                                <td width="*">VIEW</td>
                                <td width="1%">
                                    <input
                                        bind:group={adminrule_rule_field}
                                        type="checkbox"
                                        value="REPORTWINLOSE-VIEW"
                                    />
                                </td>
                                <td width="*">VIEW</td>
                            </tr>
                            <tr>
                                <td width="1%">
                                    <input
                                        bind:group={adminrule_rule_field}
                                        type="checkbox"
                                        value="PERIODE-SAVE"
                                    /></td
                                >
                                <td width="*">SAVE</td>

                                <td colspan="2">&nbsp;</td>
                                <td colspan="2">&nbsp;</td>
                            </tr>
                        </tbody>
                    </table>
                    <table class="table">
                        <thead>
                            <tr>
                                <th colspan="2">PASARAN</th>
                                <th colspan="2">ADMIN MANAGEMENT</th>
                                <th colspan="2">ADMIN RULE</th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr>
                                <td width="1%">
                                    <input
                                        bind:group={adminrule_rule_field}
                                        type="checkbox"
                                        value="PASARAN-VIEW"
                                    />
                                </td>
                                <td width="*">VIEW</td>
                                <td width="1%">
                                    <input
                                        bind:group={adminrule_rule_field}
                                        type="checkbox"
                                        value="ADMIN-VIEW"
                                    />
                                </td>
                                <td width="*">VIEW</td>
                                <td width="1%">
                                    <input
                                        bind:group={adminrule_rule_field}
                                        type="checkbox"
                                        value="ADMINRULE-VIEW"
                                    />
                                </td>
                                <td width="*">VIEW</td>
                            </tr>
                            <tr>
                                <td width="1%">
                                    <input
                                        bind:group={adminrule_rule_field}
                                        type="checkbox"
                                        value="PASARAN-SAVE"
                                    /></td
                                >
                                <td width="*">SAVE</td>
                                <td width="1%">
                                    <input
                                        bind:group={adminrule_rule_field}
                                        type="checkbox"
                                        value="ADMIN-SAVE"
                                    /></td
                                >
                                <td width="*">SAVE</td>
                                <td width="1%">
                                    <input
                                        bind:group={adminrule_rule_field}
                                        type="checkbox"
                                        value="ADMINRULE-SAVE"
                                    /></td
                                >
                                <td width="*">SAVE</td>
                            </tr>
                        </tbody>
                    </table>
                </div>
            </div>
        </Col>
    </Row>
</Container>
