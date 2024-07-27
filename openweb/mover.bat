@echo off
robocopy "D:\DB\Download" "C:\go\APP\openweb\copy" /e
set "sumber=D:\DB\Download\*"
set "tujuan=C:\go\APP\openweb\pindah\"

move "%sumber%" "%tujuan%"
if errorlevel 1 (
  echo Gagal memotong file!
) else (
  echo File berhasil dipotong ke %tujuan%
)
