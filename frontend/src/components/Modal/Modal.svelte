<script
  lang="ts"
>
  export let imgSrc: string | null

  let isOpened = false

  $: isOpened = Boolean(imgSrc)
  
  const size = 'lg'

  /**
   * Modal bezárása.
   * @param e - A kiváltó esemény.
   */
  function onModalClose (e: Event): void {
    e.preventDefault()

    imgSrc = null
  }

</script>

{#if isOpened}
  <div
    class="modal"
    tabindex="-1"
    role="dialog"
    aria-hidden={false}
  >
    <div class="modal-dialog modal-lg" role="document">
      <div class="modal-content">
        <div class="modal-header">
          <h5 class="modal-title">Kép</h5>
          <button
            type="button"
            class="btn close"
            data-dismiss="modal"
            aria-label="Close"
            on:click={ onModalClose }
          >
            <span aria-hidden="true">&times;</span>
          </button>

        </div>

        <div class="modal-body">
          {#if imgSrc}
            <img class="img-fluid" src={ imgSrc } alt="Megjelenítendő kép" />
          {/if}
        </div>

        <div class="modal-footer">
          <button
            type="button"
            class="btn btn-secondary"
            on:click={ onModalClose }
          >
            Bezárás
          </button>
        </div>
      </div>
    </div>
  </div>

  <div class="modal-backdrop show" />
{/if}

<style>
  .modal {
    display: block;
  }
</style>