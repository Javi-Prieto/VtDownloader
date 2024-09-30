import subprocess
import threading
from pathlib import Path
from subprocess import Popen

from assets.messages import SuccessMessages, ErrorMessages
from common.MOD_S_SSE.sse import receive_download_message


class DownloadsService:
    def __init__(self):
        self.downloads_path = str(Path.home() / "Downloads")
        self.aria2c_path = Path(__file__).resolve().parents[2] / 'aria2c.exe'

    def start_download_file(self, download_url: str, download_id: int, download_anyways: bool) -> (str, bool):
        if download_url.startswith("https://") or download_anyways:
            command = [self.aria2c_path, '-d', self.downloads_path, download_url]
            process = subprocess.Popen(command, stdout=subprocess.PIPE, stderr=subprocess.PIPE, text=True)
            thread = threading.Thread(target=self.download_file, args=(process,download_id))
            if process.stderr != "":
                thread.start()
                return SuccessMessages.download_file_started_message, False
            else:
                return ErrorMessages.download_file_error_message, True
        else: return ErrorMessages.not_valid_url_file_error, True

    @staticmethod
    def download_file(process: Popen, download_id: int):
        return_code = process.wait()
        if return_code == 0:
            receive_download_message(download_id, "success")
        else:
            receive_download_message(download_id,"error")
