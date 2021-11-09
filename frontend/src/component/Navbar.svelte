<script>
    import Modal from "../component/Modal.svelte"
    let myModal = "";
    const handlePopup = () => {
        myModal = new bootstrap.Modal(document.getElementById("modalentry"));
        myModal.show();
    };
</script>
<nav class="navbar fixed-top navbar-light d-flex flex-wrap justify-content-center py-2 mb-5 " style="height:50px;background-color:#1e1e1e;">
    <center>
        <img style="margin-top:-5px;" src="logo_sdsbresult.png" width="100" alt="" />
    </center>    
</nav>

<Modal
	modal_id="modalentry"
	modal_size="modal-dialog-centered"
	modal_title=""
    modal_footer_css="padding:5px;"
	modal_footer={false}>
	<slot:template slot="body">
        <center>
            Authentication Required
		    <br>
		    Users must be located in Europe
        </center>
	</slot:template>
</Modal>