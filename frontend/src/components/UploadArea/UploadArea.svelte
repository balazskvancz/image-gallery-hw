<script 
  lang="ts"
>
  import { onRefreshData } from '../../store'

  import ajax from '../../ajax/ajax'

  let input: HTMLInputElement

  let file: File | null = null

  /**
   * Feltöltés gomb eseménykezelője.
   * @param e - A kiváltó esemény.
   */
  async function onClickUpload (e: Event): Promise<void> {
    e.preventDefault()

    if (!file) {
      return
    }

    const formData = new FormData()

    formData.append("image", file)

    const isSuccess = await ajax.insert(formData)

    if (isSuccess) {
      onRefreshData.set(true)
      onRefreshData.set(false)
    }
  }

  /**
   * Input módosítását figyelő eseménykezelő.
   * @param e - A kiváltó esemény.
   */
   function onChangeFileInput (e: Event): void {
    e.preventDefault()

    const { files } = input

    if (!files) {
      return
    }

    // Ha nem volt fájl kiválasztva, akkor ennyi.
    if (files.length === 0) {
      return
    }

    if (files.length > 1) {
      return
    }

    const uploadedFile = files.item(0)

    if (!uploadedFile) {
      return
    }

    file = uploadedFile
  }
</script>

<div class="col-sm-12 col-md-6 mx-auto">
  <div class="card mb-3">
    <div class="card-header">
      <h2>Új fájl feltöltése</h2>
    </div>
  
    <div class="card-body">
      <input
        on:change={ onChangeFileInput }
        bind:this={ input }
        type="file"
    />

    </div>

    <div class="card-footer text-center">
      <button 
        type='button' 
        class="btn btn-success"
        on:click={ onClickUpload }
      >
        Feltöltés
      </button>
    </div>
  </div>
</div>