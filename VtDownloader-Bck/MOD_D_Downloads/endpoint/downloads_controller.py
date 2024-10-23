from fastapi import APIRouter
from fastapi import Response, status

from MOD_D_Downloads.service.downloads_service import DownloadsService
from MOD_D_Downloads.model.downloads_model import DownloadFileRequest, DownloadFileResponse

router = APIRouter(prefix='/download')

service = DownloadsService()

@router.post(path='/file')
def download_file(download_request: DownloadFileRequest, response: Response) -> DownloadFileResponse:
    file_name = download_request.download_url.split('/')[-1]
    service_return,is_error,d_status=DownloadsService.start_download_file(service, download_request.download_url, download_request.download_id, download_request.download_anyways)
    to_return = DownloadFileResponse(status_description=service_return,download_id=download_request.download_id,download_name=file_name,status=d_status)
    DownloadFileResponse.model_validate(to_return)
    if is_error:
        response.status_code = status.HTTP_406_NOT_ACCEPTABLE
        return to_return
    else:
        response.status_code = status.HTTP_200_OK
        return to_return



