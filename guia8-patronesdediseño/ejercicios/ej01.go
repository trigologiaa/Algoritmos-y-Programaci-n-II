package ejercicios

//FileMp3: Estructura que no se puede modificar.
type FileMp3 struct {
	nombre string
}

//PlayMp3: Función que reproduce archivos MP3.
func (p *FileMp3) PlayMp3() string {
	return "Reproduciendo archivo MP3. Nombre: " + p.nombre
}

//NewFileMp3: Función que crea instancia de FileMp3.
func NewFileMp3(nombre string) *FileMp3 {
	return &FileMp3 {
		nombre: nombre,
	}
}

//ArchivoDeMusica: Interfaz a cumplir.
type ArchivoDeMusica interface {
	Reproducir() string
}

//archivoDeMusicaAdapter: Estructura que implementa interfaz.
type archivoDeMusicaAdapter struct {
	mp3File FileMp3
}

//NewArchivoDeMusicaAdapter: Función que crea instancia de ArchivoDeMusicaAdapter.
func NewArchivoDeMusicaAdapter(mp3File FileMp3) ArchivoDeMusica {
	return &archivoDeMusicaAdapter {
		mp3File: mp3File,
	}
}

//Reproducir: Función que implementa método Reproducir().
func (p *archivoDeMusicaAdapter) Reproducir() string {
	return p.mp3File.PlayMp3()
}