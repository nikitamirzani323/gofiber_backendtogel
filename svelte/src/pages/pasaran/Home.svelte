<script>
    import Panel from "../../components/Panel.svelte";
    import Loader from "../../components/Loader.svelte";
    import { Icon, Row, Col, Container } from "sveltestrap";
    import { createEventDispatcher } from "svelte";
    let page = "Pasaran";
    let screen_height = screen.height - 480;
    export let listPasaran = [];
    export let totalrecord = 0;
    let dispatch = createEventDispatcher();
    const RefreshHalaman = () => {
        dispatch("handleRefreshData", "call");
    };
    const EditData = (e) => {
        const pasaran = {
            e,
        };
        dispatch("handleEditData", pasaran);
    };

    let searchPasaran = ""; 
    let filterPasaran = [];
    $: {
        if (searchPasaran) {
            filterPasaran = listPasaran.filter(
                (item) =>
                    item.nmpasarantogel
                        .toLowerCase()
                        .includes(searchPasaran.toLowerCase()) 
            );
        } else {
            filterPasaran = [...listPasaran];
        }
    }
</script>

<Container style="margin-top: 70px;">
    <Row>
        <Col>
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
                        <div class="row g-2">
                            <div class="col-sm-11">
                                <input
                                    bind:value={searchPasaran}
                                    type="text"
                                    class="form-control"
                                    placeholder="Search"
                                    aria-label="Search"
                                />
                            </div>
                            <div class="col-sm">
                                <button
                                    class="btn btn-warning"
                                    style="border-radius: 0px;"
                                >
                                    Search
                                </button>
                            </div>
                        </div>
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
                                        >STATUS</th
                                    >
                                    <th
                                        width="1%"
                                        style="text-align: center;vertical-align:top;font-size: 14px;"
                                        >&nbsp;</th
                                    >
                                    <th
                                        width="*"
                                        style="text-align: left;vertical-align:top;font-size: 14px;"
                                        >PASARAN</th
                                    >
                                    <th
                                        width="15%"
                                        style="text-align: left;vertical-align:top;font-size: 14px;"
                                        >HARI DIUNDI</th
                                    >
                                    <th
                                        width="10%"
                                        style="text-align: center;vertical-align:top;font-size: 14px;"
                                        >TUTUP</th
                                    >
                                    <th
                                        width="10%"
                                        style="text-align: center;vertical-align:top;font-size: 14px;"
                                        >JADWAL</th
                                    >
                                    <th
                                        width="10%"
                                        style="text-align: center;vertical-align:top;font-size: 14px;"
                                        >OPEN</th
                                    >
                                    <th
                                        width="10%"
                                        style="text-align: right;vertical-align:top;font-size: 14px;"
                                        >DISPLAY</th
                                    >
                                </tr>
                            </thead>
                            <tbody>
                                {#each filterPasaran as rec}
                                    <tr>
                                        <td
                                            on:click={() => {
                                                EditData(rec.idcomppasaran);
                                            }}
                                            NOWRAP
                                            style="text-align: center;vertical-align:top;font-size: 13px;cursor:pointer;"
                                        >
                                            <Icon name="pencil" />
                                        </td>
                                        <td
                                            NOWRAP
                                            style="text-align: center;vertical-align:top;font-size: 13px;{rec.statuspasaran_css}"
                                            >{rec.statuspasaran}</td
                                        >
                                        <td
                                            NOWRAP
                                            style="text-align: center;vertical-align:top;font-size: 13px;{rec.statuspasaranactive_css}"
                                            >{rec.statuspasaranactive}</td
                                        >
                                        <td
                                            NOWRAP
                                            style="text-align: left;vertical-align:top;font-size: 13px;"
                                            >{rec.nmpasarantogel}</td
                                        >
                                        <td
                                            NOWRAP
                                            style="text-align: left;vertical-align:top;font-size: 13px;"
                                            >{rec.pasarandiundi}</td
                                        >
                                        <td
                                            NOWRAP
                                            style="text-align: center;vertical-align:top;font-size: 13px;"
                                            >{rec.jamtutup}</td
                                        >
                                        <td
                                            NOWRAP
                                            style="text-align: center;vertical-align:top;font-size: 13px;"
                                            >{rec.jamjadwal}</td
                                        >
                                        <td
                                            NOWRAP
                                            style="text-align: center;vertical-align:top;font-size: 13px;"
                                            >{rec.jamopen}</td
                                        >
                                        <td
                                            NOWRAP
                                            style="text-align: right;vertical-align:top;font-size: 13px;"
                                            >{rec.displaypasaran}</td
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
