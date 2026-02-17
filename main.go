package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	router "github.com/v2fly/v2ray-core/v5/app/router/routercommon"
	"google.golang.org/protobuf/proto"
)

func main() {

	fileFlag := flag.String("file", "", "Path to local .dat file")
	urlFlag := flag.String("url", "", "URL to fetch .dat file from")
	outDir := flag.String("out", "out", "Output directory for files")
	fileType := flag.String("type", "geosite", "Type of .dat file: 'geosite' or 'geoip'")
	geoSiteFile := flag.String("geosite", "", "Path to GeoSite .dat file")
	geoIPFile := flag.String("geoip", "", "Path to GeoIP .dat file")
	geoSiteFiles := flag.String("geosites", "", "Comma-separated paths to GeoSite .dat files")
	geoIPFiles := flag.String("geoips", "", "Comma-separated paths to GeoIP .dat files")
	geoSiteURLs := flag.String("geosite-urls", "", "Comma-separated URLs to GeoSite .dat files")
	geoIPURLs := flag.String("geoip-urls", "", "Comma-separated URLs to GeoIP .dat files")
	flag.Parse()

	// 🚫 Remove existing output directory (if it exists)
	if err := os.RemoveAll(*outDir); err != nil {
		fmt.Println("Error removing output directory:", err)
		os.Exit(1)
	}

	// ✅ Recreate output directory
	if err := os.MkdirAll(*outDir, 0755); err != nil {
		fmt.Println("Error creating output directory:", err)
		os.Exit(1)
	}

	// Handle multiple file processing
	if *geoSiteFile != "" || *geoIPFile != "" || *geoSiteFiles != "" || *geoIPFiles != "" || *geoSiteURLs != "" || *geoIPURLs != "" {

		// Process single GeoSite file
		if *geoSiteFile != "" {
			if err := processFileOrURL(*geoSiteFile, "", *outDir, "geosite"); err != nil {
				fmt.Printf("Error processing GeoSite file %s: %v\n", *geoSiteFile, err)
				os.Exit(1)
			}
		}

		// Process single GeoIP file
		if *geoIPFile != "" {
			if err := processFileOrURL(*geoIPFile, "", *outDir, "geoip"); err != nil {
				fmt.Printf("Error processing GeoIP file %s: %v\n", *geoIPFile, err)
				os.Exit(1)
			}
		}

		// Process multiple GeoSite files
		if *geoSiteFiles != "" {
			files := strings.Split(*geoSiteFiles, ",")
			for _, file := range files {
				file = strings.TrimSpace(file)
				if file != "" {
					if err := processFileOrURL(file, "", *outDir, "geosite"); err != nil {
						fmt.Printf("Error processing GeoSite file %s: %v\n", file, err)
						os.Exit(1)
					}
				}
			}
		}

		// Process multiple GeoIP files
		if *geoIPFiles != "" {
			files := strings.Split(*geoIPFiles, ",")
			for _, file := range files {
				file = strings.TrimSpace(file)
				if file != "" {
					if err := processFileOrURL(file, "", *outDir, "geoip"); err != nil {
						fmt.Printf("Error processing GeoIP file %s: %v\n", file, err)
						os.Exit(1)
					}
				}
			}
		}

		// Process multiple GeoSite URLs
		if *geoSiteURLs != "" {
			urls := strings.Split(*geoSiteURLs, ",")
			for _, url := range urls {
				url = strings.TrimSpace(url)
				if url != "" {
					if err := processFileOrURL("", url, *outDir, "geosite"); err != nil {
						fmt.Printf("Error processing GeoSite URL %s: %v\n", url, err)
						os.Exit(1)
					}
				}
			}
		}

		// Process multiple GeoIP URLs
		if *geoIPURLs != "" {
			urls := strings.Split(*geoIPURLs, ",")
			for _, url := range urls {
				url = strings.TrimSpace(url)
				if url != "" {
					if err := processFileOrURL("", url, *outDir, "geoip"); err != nil {
						fmt.Printf("Error processing GeoIP URL %s: %v\n", url, err)
						os.Exit(1)
					}
				}
			}
		}

		return
	}

	// Legacy single file processing
	var data []byte
	var err error

	switch {
	case *fileFlag != "":
		data, err = os.ReadFile(*fileFlag)
		if err != nil {
			fmt.Println("Error reading file:", err)
			os.Exit(1)
		}
	case *urlFlag != "":
		resp, err := http.Get(*urlFlag)
		if err != nil {
			fmt.Println("Error fetching from URL:", err)
			os.Exit(1)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			fmt.Println("HTTP error:", resp.Status)
			os.Exit(1)
		}

		data, err = io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			os.Exit(1)
		}
	default:
		fmt.Println("Please provide either:")
		fmt.Println("  -file <path> -type <geosite|geoip>    (single file)")
		fmt.Println("  -url <url> -type <geosite|geoip>      (single file from URL)")
		fmt.Println("  -geosite <path> -geoip <path>         (single files)")
		fmt.Println("  -geosites <path1,path2,...>           (multiple GeoSite files)")
		fmt.Println("  -geoips <path1,path2,...>             (multiple GeoIP files)")
		fmt.Println("  -geosite-urls <url1,url2,...>         (multiple GeoSite URLs)")
		fmt.Println("  -geoip-urls <url1,url2,...>           (multiple GeoIP URLs)")
		os.Exit(1)
	}

	// Process based on file type
	switch *fileType {
	case "geosite":
		err = processGeoSite(data, *outDir)
	case "geoip":
		err = processGeoIP(data, *outDir)
	default:
		fmt.Printf("Unknown file type: %s. Use 'geosite' or 'geoip'\n", *fileType)
		os.Exit(1)
	}

	if err != nil {
		fmt.Printf("Error processing file: %v\n", err)
		os.Exit(1)
	}
}

func processGeoSite(data []byte, outDir string) error {
	list := new(router.GeoSiteList)
	if err := proto.Unmarshal(data, list); err != nil {
		return fmt.Errorf("failed to unmarshal as GeoSiteList: %v", err)
	}

	fmt.Printf("Processing GeoSite data with %d entries\n", len(list.Entry))

	for _, entry := range list.Entry {
		filename := filepath.Join(outDir, strings.ToLower(entry.CountryCode)+".list")
		file, err := os.Create(filename)
		if err != nil {
			fmt.Printf("Error creating file %s: %v\n", filename, err)
			continue
		}
		defer file.Close()

		for _, domain := range entry.Domain {
			prefix := ""
			switch domain.Type {
			case router.Domain_Plain:
				prefix = "DOMAIN-KEYWORD"
			case router.Domain_RootDomain:
				prefix = "DOMAIN-SUFFIX"
			case router.Domain_Regex:
				prefix = "DOMAIN-REGEX"
			case router.Domain_Full:
				prefix = "DOMAIN"
			}
			fmt.Fprintf(file, "%s,%s\n", prefix, domain.Value)
		}
	}

	return nil
}

func processGeoIP(data []byte, outDir string) error {
	var geoipList router.GeoIPList
	if err := proto.Unmarshal(data, &geoipList); err != nil {
		return fmt.Errorf("failed to unmarshal as GeoIPList: %v", err)
	}

	fmt.Printf("Processing GeoIP data with %d entries\n", len(geoipList.Entry))

	for _, geoip := range geoipList.Entry {
		filename := filepath.Join(outDir, strings.ToLower(geoip.CountryCode)+"@ipcidr.list")
		file, err := os.Create(filename)
		if err != nil {
			fmt.Printf("Error creating file %s: %v\n", filename, err)
			continue
		}
		defer file.Close()

		for _, cidr := range geoip.Cidr {
			ipStr := net.IP(cidr.GetIp()).String()
			prefix := cidr.GetPrefix()
			fmt.Fprintf(file, "IP-CIDR,%s/%d\n", ipStr, prefix)
		}

		fmt.Printf("Processed %d CIDR blocks for %s\n", len(geoip.Cidr), geoip.CountryCode)
	}

	return nil
}

func processFileOrURL(filePath, url, outDir, fileType string) error {
	var data []byte
	var err error
	var source string

	if filePath != "" {
		source = filePath
		fmt.Printf("Processing %s file: %s\n", fileType, filePath)
		data, err = os.ReadFile(filePath)
		if err != nil {
			return fmt.Errorf("error reading file: %v", err)
		}
	} else if url != "" {
		source = url
		fmt.Printf("Processing %s URL: %s\n", fileType, url)
		resp, err := http.Get(url)
		if err != nil {
			return fmt.Errorf("error fetching from URL: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("HTTP error: %s", resp.Status)
		}

		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("error reading response body: %v", err)
		}
	} else {
		return fmt.Errorf("either filePath or url must be provided")
	}

	switch fileType {
	case "geosite":
		if err := processGeoSite(data, outDir); err != nil {
			return fmt.Errorf("error processing %s as GeoSite: %v", source, err)
		}
	case "geoip":
		if err := processGeoIP(data, outDir); err != nil {
			return fmt.Errorf("error processing %s as GeoIP: %v", source, err)
		}
	default:
		return fmt.Errorf("unknown file type: %s", fileType)
	}

	return nil
}

// WriteGeoSiteDAT creates a .dat file from domain rules
func WriteGeoSiteDAT(entries []*router.GeoSite, filename string) error {
	geoSiteList := &router.GeoSiteList{
		Entry: entries,
	}

	data, err := proto.Marshal(geoSiteList)
	if err != nil {
		return fmt.Errorf("failed to marshal GeoSiteList: %v", err)
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write file: %v", err)
	}

	return nil
}

// WriteGeoIPDAT creates a .dat file from IP CIDR rules
func WriteGeoIPDAT(entries []*router.GeoIP, filename string) error {
	geoIPList := &router.GeoIPList{
		Entry: entries,
	}

	data, err := proto.Marshal(geoIPList)
	if err != nil {
		return fmt.Errorf("failed to marshal GeoIPList: %v", err)
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write file: %v", err)
	}

	return nil
}
