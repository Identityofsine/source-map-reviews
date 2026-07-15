import { computed, inject, Injectable, Service } from '@angular/core';
import { rxResource } from '@angular/core/rxjs-interop';
import { LookupsService as LKService } from '@arch-shared/data-source';

@Service()
export class LookupsService {

  readonly lookupsHttpService = inject(LKService)

  readonly lookupsRequest = rxResource({
    stream: () => this.lookupsHttpService.getLookups(),
  });

  private readonly _lookups = computed(() => this.lookupsRequest.value()?.lookups ?? {});

}

