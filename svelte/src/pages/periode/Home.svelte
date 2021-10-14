<script>
    import Panel from "../../components/Panel.svelte";
    import Loader from "../../components/Loader.svelte";
    import { Icon, Row, Col, Container } from "sveltestrap";
    import { createEventDispatcher } from "svelte";
    import Modal from "../../components/Modal.svelte";
    let page = "Periode";
    let sData = "New";
    let screen_height = screen.height - 480;
    export let listPeriode = [];
    export let listPeriodePasaran = [];
    export let totalrecord = 0;
    export let token = "";

    let select_pasaran = "";
    let css_loader = "display: none;";
    let msgloader = "";
    let dispatch = createEventDispatcher();
    const RefreshHalaman = () => {
        dispatch("handleRefreshData", "call");
    };
    const NewData = () => {
        sData = "New";
        let myModal = new bootstrap.Modal(document.getElementById("modalNew"));
        myModal.show();
    };
    const EditData = (e, f) => {
        const pasaran = {
            e,
            f,
        };
        dispatch("handleEditData", pasaran);
    };

    let searchPasaran = ""; 
    let filterPasaran = [];
    $: {
        if (searchPasaran) {
            filterPasaran = listPeriode.filter(
                (item) =>
                    item.pasaran_name
                        .toLowerCase()
                        .includes(searchPasaran.toLowerCase()) ||
                    item.pasaran_status
                        .toLowerCase()
                        .includes(searchPasaran.toLowerCase())
            );
        } else {
            filterPasaran = [...listPeriode];
        }
    }
    async function SaveTransaksi() {
        let flag = false;
        let msg = "";
        css_loader = "display: inline-block;";
        msgloader = "Sending...";
        if (select_pasaran == "") {
            flag = true;
            msg += "The Pasaran is required\n";
        }
        if (flag == false) {
            const res = await fetch("/api/savepasarannew", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: sData,
                    pasaran_code: select_pasaran,
                }),
            });
            const json = await res.json();

            if (json.status == 200) {
                msgloader = json.message;
            } else {
                msgloader = json.message;
            }
            setTimeout(function () {
                css_loader = "display: none;";
            }, 1000);
            RefreshHalaman();

            select_pasaran = "";
        } else {
            alert(msg);
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
                            bind:value={searchPasaran}
                            type="text"
                            class="form-control"
                            placeholder="Search Pasaran"
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
                                        width="1%"
                                        style="text-align: center;vertical-align:top;font-size: 14px;"
                                        >&nbsp;</th
                                    >
                                    <th
                                        width="10%"
                                        style="text-align: center;vertical-align:top;font-size: 14px;"
                                        >DATE</th
                                    >
                                    <th
                                        NOWRAP
                                        width="1%"
                                        style="text-align: left;vertical-align:top;font-size: 14px;"
                                        >INVOICE</th
                                    >
                                    <th
                                        width="10%"
                                        style="text-align: left;vertical-align:top;font-size: 14px;"
                                        >PERIODE</th
                                    >
                                    <th
                                        width="*"
                                        style="text-align: left;vertical-align:top;font-size: 14px;"
                                        >PASARAN</th
                                    >
                                    <th
                                        width="10%"
                                        style="text-align: center;vertical-align:top;font-size: 14px;"
                                        >KELUARAN</th
                                    >
                                    <th
                                        NOWRAP
                                        width="1%"
                                        style="text-align: right;vertical-align:top;font-size: 14px;"
                                        >REVISI</th
                                    >
                                    <th
                                        width="10%"
                                        style="text-align: right;vertical-align:top;font-size: 14px;"
                                        >MEMBER</th
                                    >
                                    <th
                                        width="10%"
                                        style="text-align: right;vertical-align:top;font-size: 14px;"
                                        >BET</th
                                    >
                                    <th
                                        width="10%"
                                        style="text-align: right;vertical-align:top;font-size: 14px;"
                                        >BAYAR</th
                                    >
                                    <th
                                        width="10%"
                                        style="text-align: right;vertical-align:top;font-size: 14px;"
                                        >WINLOSE</th
                                    >
                                </tr>
                            </thead>
                            <tbody>
                                {#each filterPasaran as rec}
                                    <tr>
                                        <td
                                            on:click={() => {
                                                EditData(
                                                    rec.pasaran_invoice,
                                                    rec.pasaran_code
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
                                            >{rec.no}</td
                                        >
                                        <td
                                            NOWRAP
                                            style="text-align: center;vertical-align:top;font-size: 13px;{rec.pasaran_status_css}"
                                            >{rec.pasaran_status}</td
                                        >

                                        <td
                                            NOWRAP
                                            style="text-align: center;vertical-align:top;font-size: 13px;"
                                            >{rec.pasaran_tanggal}</td
                                        >
                                        <td
                                            NOWRAP
                                            style="text-align: left;vertical-align:top;font-size: 13px;"
                                            >{rec.pasaran_invoice}</td
                                        >
                                        <td
                                            NOWRAP
                                            style="text-align: left;vertical-align:top;font-size: 13px;"
                                            >{rec.pasaran_periode}</td
                                        >
                                        <td
                                            NOWRAP
                                            style="text-align: left;vertical-align:top;font-size: 13px;"
                                            >{rec.pasaran_name}</td
                                        >
                                        <td
                                            NOWRAP
                                            style="text-align: center;vertical-align:top;font-size: 13px;font-weight:bold;color:black;"
                                            >{rec.pasaran_keluaran}</td
                                        >
                                        <td
                                            NOWRAP
                                            style="text-align: right;vertical-align:top;font-size: 13px;{rec.pasaran_revisi_css}"
                                            >{new Intl.NumberFormat().format(
                                                rec.pasaran_revisi
                                            )}
                                            {#if rec.pasaran_revisi > 0} 
                                            <i class="bi bi-chat-right-text" data-bs-toggle="tooltip" data-bs-placement="top" title="{rec.pasaran_msgrevisi}"></i> 
                                            {/if}
                                        </td>
                                        <td
                                            NOWRAP
                                            style="text-align: right;vertical-align:top;font-size: 13px;{rec.pasaran_totalmember_css}"
                                            >{new Intl.NumberFormat().format(
                                                rec.pasaran_totalmember
                                            )}</td
                                        >
                                        <td
                                            NOWRAP
                                            style="text-align: right;vertical-align:top;font-size: 13px;{rec.pasaran_totalbet_css}"
                                            >{new Intl.NumberFormat().format(
                                                rec.pasaran_totalbet
                                            )}</td
                                        >
                                        <td
                                            NOWRAP
                                            style="text-align: right;vertical-align:top;font-size: 13px;{rec.pasaran_totaloutstanding_css}"
                                            >{new Intl.NumberFormat().format(
                                                rec.pasaran_totaloutstanding
                                            )}</td
                                        >
                                        <td
                                            NOWRAP
                                            style="text-align: right;vertical-align:top;font-size: 13px;{rec.pasaran_winlose_css}"
                                            >{new Intl.NumberFormat().format(
                                                rec.pasaran_winlose
                                            )}</td
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
    modal_body_height={"height:100px;"}
    modal_footer_flag={true}
>
    <slot:template slot="header">
        <h5 class="modal-title" id="exampleModalLabel">Entry/{sData}</h5>
    </slot:template>
    <slot:template slot="body">
        <Container>
            <Row>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Pasaran</label>
                    <select
                        class="form-control required"
                        bind:value={select_pasaran}
                    >
                        {#each listPeriodePasaran as rec}
                            <option value={rec.pasarancomp_idcompp}
                                >{rec.pasarancomp_nama}</option
                            >
                        {/each}
                    </select>
                </div>
            </Row>
        </Container>
    </slot:template>
    <slot:template slot="footer">
        <div class="float-end">
            <button
                on:click={() => {
                    SaveTransaksi();
                }}
                class="btn btn-warning"
                style="border-radius: 0px;"
            >
                Save
            </button>
        </div>
    </slot:template>
</Modal>
