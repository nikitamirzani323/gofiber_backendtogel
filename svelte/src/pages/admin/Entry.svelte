<script>
    import { Icon } from "sveltestrap";
    import { createEventDispatcher } from "svelte";

    export let sData = "";
    export let token = "";
    export let admin_username = "";
    export let admin_listrule = [];
    export let admin_listip = [];
    export let admin_name_field = "";
    export let admin_type_field = "";
    export let admin_idrule_field = 0;
    export let admin_status_field = "";
    export let admin_create_field = "";
    export let admin_update_field = "";
    let admin_password = "";
    let admin_ipaddress = "";
    let css_loader = "display: none;";
    let css_panel_iplist = "display: none;";
    let msgloader = "";
    let dispatch = createEventDispatcher();
    const BackHalaman = () => {
        dispatch("handleBackHalaman", "call");
    };
    async function SaveTransaksi() {
        let flag = false;
        let msg = "";
        css_loader = "display: inline-block;";
        msgloader = "Sending...";
        if (admin_username == "") {
            flag = true;
            msg += "The Username is required\n";
        }
        if (admin_name_field == "") {
            flag = true;
            msg += "The Name is required\n";
        }
        if (admin_type_field == "MASTER") {
            select_listrule = 0;
        }

        if (flag == false) {
            const res = await fetch("/api/saveadmin", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: sData,
                    idruleadmin: parseInt(admin_idrule_field),
                    page: "ADMIN-SAVE",
                    username: admin_username,
                    password: admin_password,
                    nama: admin_name_field,
                    status: admin_status_field,
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
    async function saveNewIpList() {
        let flag = false;
        let msg = "";
        let totaliplist = admin_listip.length;
        css_loader = "display: inline-block;";
        msgloader = "Sending...";
        if (admin_ipaddress == "") {
            flag = true;
            msg += "The Ipaddress is required";
        }
        if (parseInt(totaliplist) > 4) {
            flag = true;
            msg += "Maximal 5 Ipaddress";
        }
        if (flag == false) {
            const res = await fetch("/api/saveadminiplist", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sData: "New",
                    page: "ADMIN-SAVE",
                    username: admin_username,
                    ipaddress: admin_ipaddress,
                }),
            });
            const json = await res.json();

            if (json.status == 200) {
                msgloader = json.message;

                dispatch("handleRefreshIplist", admin_username);
                admin_ipaddress = "";
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
            admin_ipaddress = "";
            setTimeout(function () {
                css_loader = "display: none;";
            }, 1000);
        }
    }
    async function DeleteIpList(e) {
        const iplist = {
            e,
            admin_username,
        };
        dispatch("handleDeleteIplist", iplist);
    }
    async function NewIplist() {
        let myModal = new bootstrap.Modal(
            document.getElementById("modalNewIpaddress")
        );
        myModal.show();
    }
    if(admin_type_field != "MASTER") {
        css_panel_iplist = "";
    }
</script>

<div id="loader" style="margin-left:50%;{css_loader}">
    {msgloader}
</div>
<div class="container-fluid" style="margin-top: 70px;">
    <div class="row">
        <div class="col-sm-12">
            <button
                on:click={() => {
                    BackHalaman();
                }}
                class="btn btn-dark btn-sm"
                style="border-radius: 0px;">
                Back
            </button>
        </div>
        <div class="clearfix" />
        <div class="col-sm-3">
            <div class="card" style="border-radius: 0px;margin-top:10px;">
                <div class="card-header" style="">
                    Pasaran / {sData}
                    <div class="float-end">
                        <button
                            on:click={() => {
                                SaveTransaksi();
                            }}
                            class="btn btn-warning btn-sm"
                            style="border-radius: 0px;">
                            Save
                        </button>
                    </div>
                </div>
                <div class="card-body" style="height:500px;">
                    <div class="mb-3">
                        <label for="exampleForm" class="form-label">Username</label>
                        <input
                            bind:value={admin_username}
                            type="text"
                            disabled
                            maxlength="70"
                            class="form-control"
                            placeholder="Username"
                            aria-label="Username"/>
                    </div>
                    <div class="mb-3">
                        <label for="exampleForm" class="form-label">Password</label>
                        <input
                            bind:value={admin_password}
                            type="password"
                            class="form-control"
                            placeholder="Password"
                            aria-label="Password"/>
                    </div>
                    {#if admin_type_field != "MASTER"}
                        <div class="mb-3">
                            <label for="exampleForm" class="form-label">Admin Rule</label>
                            <select
                                bind:value={admin_idrule_field}
                                class="form-control">
                                <option value="0">--Select--</option>
                                {#each admin_listrule as rec}
                                    <option value={rec.adminrule_idruleadmin}>{rec.adminrule_name}</option>
                                {/each}
                            </select>
                        </div>
                    {/if}
                    <div class="mb-3">
                        <label for="exampleForm" class="form-label">Name</label>
                        <input
                            bind:value={admin_name_field}
                            type="text"
                            maxlength="70"
                            class="form-control"
                            placeholder="Name"
                            aria-label="Name"/>
                    </div>
                    <div class="mb-3">
                        <label for="exampleForm" class="form-label">Status</label>
                        <select
                            class="form-control"
                            bind:value={admin_status_field}>
                            <option value="ACTIVE">ACTIVE</option>
                            <option value="BANNED">BANNED</option>
                        </select>
                    </div>
                    <div class="mb-3">
                        <table>
                            <tr>
                                <td style="font-size: 11px;">Create</td>
                                <td style="font-size: 11px;">:</td>
                                <td style="font-size: 11px;">{admin_create_field}</td>
                            </tr>
                            <tr>
                                <td style="font-size: 11px;">Update</td>
                                <td style="font-size: 11px;">:</td>
                                <td style="font-size: 11px;">{admin_update_field}</td>
                            </tr>
                        </table>
                    </div>
                </div>
            </div>
        </div>
        <div class="col-sm-3">
            <div class="card" style="border-radius: 0px;margin-top:10px;{css_panel_iplist}">
                <div class="card-header" style="">
                    IP LIST
                    <div class="float-end">
                        <button
                            on:click={() => {
                                NewIplist();
                            }}
                            class="btn btn-primary btn-sm">
                            New
                        </button>
                    </div>
                </div>
                <div class="card-body" style="height:455px;">
                    <table class="table table-striped">
                        <thead>
                            <tr>
                                <th
                                    width="1%"
                                    style="text-align:center;font-size:14px;">&nbsp;</th>
                                <th
                                    width="1%"
                                    style="text-align:center;font-size:14px;">NO</th>
                                <th
                                    width="*"
                                    style="text-align:left;font-size:14px;">IPADDRESS</th>
                            </tr>
                        </thead>
                        <tbody>
                            {#each admin_listip as rec}
                                <tr>
                                    <td
                                        on:click={() => {
                                            DeleteIpList(
                                                rec.adminiplist_idcompiplist
                                            );
                                        }}
                                        style="text-align:center;font-size:12px;cursor:pointer;"><Icon name="trash" /></td>
                                    <td
                                        style="text-align:center;font-size:12px;">{rec.adminiplist_no}</td>
                                    <td
                                        style="text-align:left;font-size:12px;">{rec.adminiplist_iplist}</td>
                                </tr>
                            {/each}
                        </tbody>
                    </table>
                </div>
                <div class="card-footer" style="margin:0px;padding:10px;font-size:13px;">
                    <b>TOTAL RECORD : {admin_listip.length}</b>
                </div>
            </div>
        </div>
    </div>
</div>
<div
    class="modal fade"
    id="modalNewIpaddress"
    tabindex="-1"
    aria-labelledby="exampleModalLabel"
    aria-hidden="true">
    <div class="modal-dialog">
        <div class="modal-content">
            <div class="modal-header">
                <h5 class="modal-title" id="exampleModalLabel">
                    New IPADDRESS
                </h5>
                <button
                    type="button"
                    class="btn-close"
                    data-bs-dismiss="modal"
                    aria-label="Close"
                />
            </div>
            <div class="modal-body">
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">IPADDRESS</label
                    >
                    <input
                        bind:value={admin_ipaddress}
                        type="text"
                        maxlength="20"
                        class="form-control required"
                        placeholder="IPADDRESS"
                        aria-label="IPADDRESS"
                    />
                </div>
            </div>
            <div class="modal-footer">
                <button
                    on:click={() => {
                        saveNewIpList();
                    }}
                    type="button"
                    class="btn btn-warning">SAVE</button
                >
            </div>
        </div>
    </div>
</div>
