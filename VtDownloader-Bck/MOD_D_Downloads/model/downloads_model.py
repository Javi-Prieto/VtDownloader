from pydantic import BaseModel, Field, ConfigDict


class DownloadFileRequest(BaseModel):
    model_config = ConfigDict(from_attributes=True)

    download_url: str = Field(title='The url to the file to be downloaded.')
    download_id: int = Field(title='Is the id of the download.', gt=0)
    download_anyways: bool = Field(title='(Not Required) Tries to download without caring of the url.', default=False)

class DownloadFileResponse(BaseModel):
    model_config = ConfigDict(from_attributes=True)

    status_description: str
