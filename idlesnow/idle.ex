defmodule Heart do
  	def print(lang) do
		IO.puts("Hello!"<>lang)
	end
	def notify(mess) do
		{:ok, file} = File.open("../pot/broadcasts", [:read, :write])
		{:ok, oldcontents } = File.read("../pot/broadcasts")
		IO.binwrite(file, oldcontents <> mess)
		File.close(file)
	end
	def see do
		File.read("../pot/broadcasts")
	end

end
