<script 
  lang="ts"
>
  import { onDestroy, onMount } from 'svelte'

  import { onRefreshData } from '../../store'

  import ajax from '../../ajax/ajax'

  import type { TImages, TOrderBy, TOrderDirection } from '../../definitions'

  let data: TImages = []

  let orderBy: string     = 'name'
  let direction: string   = 'asc'

  /** Adatok szervertől való lekérdezése. */
  async function getData (): Promise<void> {
    data = await ajax.get(orderBy as TOrderBy, direction as TOrderDirection)
  }

  /** Adatok lekérdezését hallgató eseménykezelő. */
  const unsubscribeOnRefresh = onRefreshData.subscribe(async (needToRefresh) => {
    console.log(needToRefresh)

    if (needToRefresh) {
      await getData()
    }
  }) 

  /**
   * Törlés gomb closure-je.
   * @param name - A törlendő kép neve.
   */
  function onClickDelete (name: string): (e: Event) => Promise<void> {
    return async (e: Event): Promise<void> => {
      e.preventDefault()

      const isSuccess = await ajax.deleteByName(name)

      if (isSuccess) {
        await getData()
      }
    }
  }

  onMount(async () => {
    await getData()
  })

  onDestroy(() => {
    unsubscribeOnRefresh()
  })
</script>

<div class="container-fluid">
  <div class="row">
    <div class="col-sm-12 col-md-4">
      <label class="fw-bold">Miszerinti rendezés</label>
      <select class="form-select" bind:value={ orderBy }>
        <option value="name">Név</option>
        <option value="date">Dátum</option>
      </select>
    </div>

    <div class="col-sm-12 col-md-4">
      <label class="fw-bold">Rendezés iránya</label>
      <select class="form-select" bind:value={ direction }>
        <option value="asc">Növekvő</option>
        <option value="desc">Csökkenő</option>
      </select>
    </div>

    <div class="col-sm-12 mt-3">
      <button class="btn btn-primary" on:click={ getData }>Rendezés</button>
    </div>
    <hr class="my-3" />
  </div>

  <div class="table-responsive">
    <table class="table table-striped table-hover">
      <thead>
        <tr>
          <th class="align-middle">Név</th>
          <th class="align-middle">Feltöltés dátuma</th>
          <th class="align-middle">Megnyitás</th>
          <th class="align-middle">Törlés</th>
        </tr>
      </thead>

      <tbody>
        {#each data as image (image.name)}
          <tr>
            <td class="align-middle">{ image.name }</td>
            <td class="align-middle">{ image.createdAt }</td>
            <td class="align-middle">
              <button class="btn btn-primary">Megnyitás</button>
            </td>
            <td class="align-middle">
              <button class="btn btn-danger" on:click={ onClickDelete(image.name)}>Törlés</button>
            </td>
          </tr>
        {/each}
      </tbody>
    </table>
  </div>
</div>