<script>
    import Panel from "../../components/Panel.svelte";
    import Loader from "../../components/Loader.svelte";
    import { Icon, Row, Col, Container } from "sveltestrap";
    import { createEventDispatcher } from "svelte";
    import Modal from "../../components/Modal.svelte";
    import { createForm } from "svelte-forms-lib";
    import * as yup from "yup";
    let page = "Admin Rule Management";
    let sData = "New";
    let screen_height = screen.height - 480;
    export let token = "";
    export let listAdminrule = [];
    export let totalrecord = 0;

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
            admin_name_field: "",
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
    const RefreshHalaman = () => {
        dispatch("handleRefreshData", "call");
    };
    const NewData = () => {
        sData = "New";
        let myModal = new bootstrap.Modal(document.getElementById("modalNew"));
        myModal.show();
    };
    const EditData = (e, f) => {
        const adminpage = {
            e,
            f,
        };
        dispatch("handleEditData", adminpage);
    };
    async function SaveTransaksi(name) {
        css_loader = "display: inline-block;";
        msgloader = "Sending...";
        const res = await fetch("/api/saveadminrule", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                sdata: sData,
                page: "ADMINRULE-SAVE",
                idrule: 1,
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
        $form.admin_name_field = "";
        RefreshHalaman();
    }
    let searchAdminRule = "";
    let filterAdminRule = [];
    $: {
        if (searchAdminRule) {
            filterAdminRule = listAdminrule.filter((item) =>
                item.adminrule_nama
                    .toLowerCase()
                    .includes(searchAdminRule.toLowerCase())
            );
        } else {
            filterAdminRule = [...listAdminrule];
        }
    }
</script>

<div id="loader" style="margin-left:50%;{css_loader}">
    {msgloader}
</div>
<Container style="margin-top: 70px;">
    <Row>
        <Col>
            <button
                on:click={() => {
                    NewData();
                }}
                class="btn btn-primary"
                style="border-radius: 0px;"
            >
                New
            </button>
            <button
                on:click={() => {
                    RefreshHalaman();
                }}
                class="btn btn-primary"
                style="border-radius: 0px;"
            >
                Refresh
            </button>
            <Panel height_body="{screen_height}px">
                <slot:template slot="cheader">
                    {page}
                </slot:template>
                <slot:template slot="csearch">
                    <div class="col-lg-12" style="padding: 5px;">
                        <input
                            bind:value={searchAdminRule}
                            type="text"
                            class="form-control"
                            placeholder="Search"
                            aria-label="Search"
                        />
                    </div>
                </slot:template>
                <slot:template slot="cbody">
                    {#if totalrecord > 0}
                        <table class="table table-striped table-hover">
                            <thead>
                                <tr>
                                    <th
                                        width="1%"
                                        style="text-align: center;vertical-align:top;font-size: 14px;"
                                        >&nbsp;</th
                                    >
                                    <th
                                        width="1%"
                                        style="text-align: center;vertical-align:top;font-size: 14px;"
                                        >NO</th
                                    >
                                    <th
                                        width="*"
                                        style="text-align: left;vertical-align:top;font-size: 14px;"
                                        >NAMA</th
                                    >
                                </tr>
                            </thead>
                            <tbody>
                                {#each filterAdminRule as rec}
                                    <tr>
                                        <td
                                            on:click={() => {
                                                EditData(
                                                    rec.adminrule_id,
                                                    rec.adminrule_nama
                                                );
                                            }}
                                            NOWRAP
                                            style="text-align: center;vertical-align:top;font-size: 13px;cursor:pointer;"
                                        >
                                            <Icon name="pencil" />
                                        </td>
                                        <td
                                            NOWRAP
                                            style="text-align: center;vertical-align:top;font-size: 13px;"
                                            >{rec.adminrule_no}</td
                                        >
                                        <td
                                            NOWRAP
                                            style="text-align: left;vertical-align:top;font-size: 13px;"
                                            >{rec.adminrule_nama}</td
                                        >
                                    </tr>
                                {/each}
                            </tbody>
                        </table>
                    {:else}
                        <center>
                            <Loader />
                        </center>
                    {/if}
                </slot:template>
                <slot:template slot="cfooter">
                    Total Record : {totalrecord}
                </slot:template>
            </Panel>
        </Col>
    </Row>
</Container>
<Modal
    modal_id={"modalNew"}
    modal_size={"modal-dialog-centered"}
    modal_body_height={"height:200px;"}
    modal_footer_flag={true}
>
    <slot:template slot="header">
        <h5 class="modal-title" id="exampleModalLabel">Entry/{sData}</h5>
    </slot:template>
    <slot:template slot="body">
        <Container>
            <Row>
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
            </Row>
        </Container>
    </slot:template>
    <slot:template slot="footer">
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
    </slot:template>
</Modal>
