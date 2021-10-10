<script>
    export let modal_id = "";
    export let modal_size = "";
    export let modal_body_height = "height:350px;overflow:scroll;";
    export let modal_footer_flag = true;
</script>

<div class="modal fade" id={modal_id} tabindex="-1" data-bs-backdrop="false">
    <div class="modal-dialog {modal_size}">
        <div class="modal-content">
            <div class="modal-header" style="width:100%;">
                <slot name="header" />
                <button
                    type="button"
                    class="btn-close"
                    data-bs-dismiss="modal"
                    aria-label="Close"
                />
            </div>
            <div
                class="modal-body"
                style="font-size:12px;padding:0px;margin:0px;color:black;{modal_body_height}"
            >
                <slot name="body" />
            </div>
            {#if modal_footer_flag}
                <div class="modal-footer" style="margin:0px;border-top: none;">
                    <slot name="footer" />
                </div>
            {/if}
        </div>
    </div>
</div>
