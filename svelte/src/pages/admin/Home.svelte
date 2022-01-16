<script>
    import Panel from "../../components/Panel.svelte";
    import Loader from "../../components/Loader.svelte";
    import { Icon, Row, Col, Container } from "sveltestrap";
    import { createEventDispatcher } from "svelte";
    import Modal from "../../components/Modal.svelte";
    import { createForm } from "svelte-forms-lib";
    import * as yup from "yup";

    let page = "Admin Management";
    let sData = "New";
    let screen_height = screen.height - 480;
    export let token = "";
    export let listAdmin = [];
    export let admin_listrule = [];
    export let totalrecord = 0;

    let css_loader = "display: none;";
    let msgloader = "";

    let dispatch = createEventDispatcher();

    const schema = yup.object().shape({
        admin_username: yup
            .string()
            .required("Username is Required")
            .matches(
                /^[a-zA-z0-9]+$/,
                "Username must Character A-Z or a-z or 1-9"
            )
            .min(4, "Username must be at least 4 Character")
            .max(20, "Username must be at most 4 Character"),
        admin_password: yup
            .string()
            .required("Password is Required")
            .matches(
                /^[a-zA-z0-9]+$/,
                "Password must Character A-Z or a-z or 1-9"
            )
            .min(4, "Password must be at least 4 Character")
            .max(20, "Password must be at most 4 Character"),
        admin_name_field: yup.string().required("Name is Required"),
        admin_idrule_field: yup.number().required("Name is Required"),
    });
    const { form, errors, handleChange, handleSubmit } = createForm({
        initialValues: {
            admin_username: "",
            admin_password: "",
            admin_name_field: "",
            admin_idrule_field: "0",
        },
        validationSchema: schema,
        onSubmit: (values) => {
            SaveTransaksi(
                values.admin_username,
                values.admin_password,
                values.admin_name_field,
                values.admin_idrule_field
            );
        },
    });
    $: {
        if (
            $errors.admin_username ||
            $errors.admin_password ||
            $errors.admin_name_field ||
            $errors.admin_idrule_field
        ) {
            alert(
                $errors.admin_username +
                    "\n" +
                    $errors.admin_password +
                    "\n" +
                    $errors.admin_name_field +
                    "\n" +
                    $errors.admin_idrule_field
            );
            $form.admin_username = "";
            $form.admin_password = "";
            $form.admin_name_field = "";
            $form.admin_idrule_field = "0";
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
    const EditData = (e,f) => {
        const adminpage = {
            e,f,
        };
        dispatch("handleEditData", adminpage);
    };
    async function SaveTransaksi(username, password, name, rule) {
        let flag = true;
        let msg = "";
        if (rule < 1) {
            flag = false;
            msg += "The Admin Rule is required";
        }
        if (flag) {
            const res = await fetch("/api/saveadmin", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: sData,
                    idruleadmin: parseInt(rule),
                    page: "ADMIN-SAVE",
                    username: username,
                    password: password,
                    nama: name,
                    status: "ACTIVE",
                }),
            });
            const json = await res.json();

            if (json.status == 200) {
                msgloader = json.message;
                $form.admin_username = "";
                $form.admin_password = "";
                $form.admin_name_field = "";
                $form.admin_idrule_field = "0";
            } else if (json.status == 403) {
                alert(json.message);
            } else {
                msgloader = json.message;
            }
            setTimeout(function () {
                css_loader = "display: none;";
            }, 1000);
            RefreshHalaman();
        } else {
            alert(msg);
        }
    }
    let searchAdmin = "";
    let filterAdmin = [];
    $: {
        if (searchAdmin) {
            filterAdmin = listAdmin.filter(
                (item) =>
                    item.admin_username
                        .toLowerCase()
                        .includes(searchAdmin.toLowerCase()) ||
                    item.admin_nama
                        .toLowerCase()
                        .includes(searchAdmin.toLowerCase())
            );
        } else {
            filterAdmin = [...listAdmin];
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
                class="btn btn-primary btn-sm"
                style="border-radius: 0px;">
                New
            </button>
            <button
                on:click={() => {
                    RefreshHalaman();
                }}
                class="btn btn-primary btn-sm"
                style="border-radius: 0px;">
                Refresh
            </button>
            <Panel height_body="{screen_height}px">
                <slot:template slot="cheader">
                    {page}
                </slot:template>
                <slot:template slot="csearch">
                    <div class="col-lg-12" style="padding: 5px;">
                        <input
                            bind:value={searchAdmin}
                            type="text"
                            class="form-control"
                            placeholder="Search"
                            aria-label="Search"/>
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
                                        width="1%"
                                        style="text-align: center;vertical-align:top;font-size: 14px;"
                                        >STATUS</th
                                    >
                                    <th
                                        width="10%"
                                        style="text-align: left;vertical-align:top;font-size: 14px;"
                                        >TIMEZONE</th
                                    >
                                    <th
                                        width="10%"
                                        style="text-align: left;vertical-align:top;font-size: 14px;"
                                        >IPADDRESS</th
                                    >
                                    <th
                                        width="10%"
                                        style="text-align: center;vertical-align:top;font-size: 14px;"
                                        >LASTLOGIN</th
                                    >
                                    <th
                                        width="10%"
                                        style="text-align: center;vertical-align:top;font-size: 14px;"
                                        >JOINDATE</th
                                    >
                                    <th
                                        width="10%"
                                        style="text-align: left;vertical-align:top;font-size: 14px;"
                                        >RULE</th
                                    >
                                    <th
                                        width="10%"
                                        style="text-align: left;vertical-align:top;font-size: 14px;"
                                        >USERNAME</th
                                    >
                                    <th
                                        width="*"
                                        style="text-align: left;vertical-align:top;font-size: 14px;"
                                        >NAMA</th
                                    >
                                </tr>
                            </thead>
                            <tbody>
                                {#each filterAdmin as rec}
                                    <tr>
                                        <td
                                            on:click={() => {
                                                EditData(rec.admin_username,rec.admin_tipe);
                                            }}
                                            NOWRAP
                                            style="text-align: center;vertical-align:top;font-size: 13px;cursor:pointer;"
                                        >
                                            <Icon name="pencil" />
                                        </td>
                                        <td
                                            NOWRAP
                                            style="text-align: center;vertical-align:top;font-size: 13px;"
                                            >{rec.admin_no}</td
                                        >
                                        <td
                                            NOWRAP
                                            style="text-align: center;vertical-align:top;font-size: 13px;{rec.admin_statuscss}"
                                            >{rec.admin_status}</td
                                        >
                                        <td
                                            NOWRAP
                                            style="text-align: left;vertical-align:top;font-size: 13px;"
                                            >{rec.admin_timezone}</td
                                        >
                                        <td
                                            NOWRAP
                                            style="text-align: left;vertical-align:top;font-size: 13px;"
                                            >{rec.admin_lastipaddres}</td
                                        >
                                        <td
                                            NOWRAP
                                            style="text-align: center;vertical-align:top;font-size: 13px;"
                                            >{rec.admin_lastlogin}</td
                                        >
                                        <td
                                            NOWRAP
                                            style="text-align: center;vertical-align:top;font-size: 13px;"
                                            >{rec.admin_joindate}</td
                                        >
                                        <td
                                            NOWRAP
                                            style="text-align: left;vertical-align:top;font-size: 13px;"
                                            >{rec.admin_rule}</td
                                        >
                                        <td
                                            NOWRAP
                                            style="text-align: left;vertical-align:top;font-size: 13px;"
                                            >{rec.admin_username}</td
                                        >
                                        <td
                                            NOWRAP
                                            style="text-align: left;vertical-align:top;font-size: 13px;"
                                            >{rec.admin_nama}</td
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
    modal_body_height={"height:300px;"}
    modal_footer_flag={true}
>
    <slot:template slot="header">
        <h5 class="modal-title" id="exampleModalLabel">Entry/{sData}</h5>
    </slot:template>
    <slot:template slot="body">
        <Container>
            <Row>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Username</label>
                    <input
                        on:change={handleChange}
                        bind:value={$form.admin_username}
                        invalid={$errors.admin_username.length > 0}
                        type="text"
                        maxlength="70"
                        class="form-control"
                        placeholder="Username"
                        aria-label="Username"
                    />
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Password</label>
                    <input
                        on:change={handleChange}
                        bind:value={$form.admin_password}
                        invalid={$errors.admin_password.length > 0}
                        type="password"
                        class="form-control"
                        placeholder="Password"
                        aria-label="Password"
                    />
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label"
                        >Admin Rule</label
                    >
                    <select
                        on:change={handleChange}
                        bind:value={$form.admin_idrule_field}
                        invalid={$errors.admin_idrule_field.length > 0}
                        class="form-control"
                    >
                        <option value="0">--Select--</option>
                        {#each admin_listrule as rec}
                            <option value={rec.adminrule_idruleadmin}
                                >{rec.adminrule_name}</option
                            >
                        {/each}
                    </select>
                </div>
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
