import unittest

from MOD_D_Downloads.service.downloads_service import DownloadsService
service = DownloadsService()

class TestDownloadService(unittest.TestCase):

    def test_download_start(self):
        response  = DownloadsService.start_download_file(service, "", 1, False)
        self.assertEqual("The url provided is not secure please.", response)
        response = DownloadsService.start_download_file(service, "https://download-cdn.jetbrains.com/python/pycharm-community-2024.2.2-aarch64.exe", 2, False)
        self.assertEqual("Your download started correctly.", response)

if __name__ == '__main__':
    unittest.main()
