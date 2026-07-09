import { HttpClient } from "@angular/common/http"
import { inject, Injectable } from "@angular/core"
import { UploadImageResponse } from "@arch-shared/types"
import { Observable } from "rxjs"

Injectable({
  providedIn: 'root',
})
export class ImageService {
  readonly API_URL = `/api/images`
  readonly http = inject(HttpClient);


  public uploadImage(file: File, imageRequest: {
    caption?: string;
    fileExt?: string;
  }): Observable<UploadImageResponse> {
    const formData = new FormData();
    formData.append('file', file);
    formData.append('json', JSON.stringify(imageRequest));
    return this.http.post<UploadImageResponse>(`${this.API_URL}/upload`,
      formData,
    );
  }

}
