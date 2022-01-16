<script>
    import Panel from "../../components/Panel.svelte";
    import Loader from "../../components/Loader.svelte";
    import { Row, Col, Container } from "sveltestrap";
    import { createEventDispatcher } from "svelte";

    let page = "Log Management";
    let screen_height = screen.height - 480;
    export let token = "";
    export let listHome = [];
    export let totalrecord = 0;

    let css_loader = "display: none;";
    let msgloader = "";

    let dispatch = createEventDispatcher();
    const RefreshHalaman = () => {
        dispatch("handleRefreshData", "call");
    };
    let searchHome = "";
    let filterHome = [];
    $: {
        if (searchHome) {
            filterHome = listHome.filter(
                (item) =>
                    item.log_username
                        .toLowerCase()
                        .includes(searchHome.toLowerCase()) ||
                    item.log_note
                        .toLowerCase()
                        .includes(searchHome.toLowerCase()) ||
                    item.log_page
                        .toLowerCase()
                        .includes(searchHome.toLowerCase())
            );
        } else {
            filterHome = [...listHome];
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
                            bind:value={searchHome}
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
                                        style="text-align: center;vertical-align:top;font-size: 14px;">NO</th>
                                    
                                    <th
                                        width="10%"
                                        style="text-align: center;vertical-align:top;font-size: 14px;">DATETIME</th>
                                    <th
                                        width="10%"
                                        style="text-align: left;vertical-align:top;font-size: 14px;">USERNAME</th>
                                    <th
                                        width="10%"
                                        style="text-align: left;vertical-align:top;font-size: 14px;">PAGE</th>
                                    <th
                                        width="10%"
                                        style="text-align: left;vertical-align:top;font-size: 14px;">ACTION</th>
                                    <th
                                        width="*"
                                        style="text-align: left;vertical-align:top;font-size: 14px;">NOTE</th>
                                </tr>
                            </thead>
                            <tbody>
                                {#each filterHome as rec}
                                    <tr>
                                        <td
                                            NOWRAP
                                            style="text-align: center;vertical-align:top;font-size: 13px;">{rec.log_no}</td>
                                        <td
                                            NOWRAP
                                            style="text-align: center;vertical-align:top;font-size: 13px;">{rec.log_datetime}</td>
                                        <td
                                            NOWRAP
                                            style="text-align: left;vertical-align:top;font-size: 13px;">{rec.log_username}</td>
                                        <td
                                            NOWRAP
                                            style="text-align: left;vertical-align:top;font-size: 13px;">{rec.log_page}</td>
                                        <td
                                            NOWRAP
                                            style="text-align: left;vertical-align:top;font-size: 13px;">{rec.log_tipe}</td>
                                        <td
                                            style="text-align: left;vertical-align:top;font-size: 13px;">{@html rec.log_note}</td>
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

