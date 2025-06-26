package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

var GenSwaggerCmd = &cobra.Command{
	Use:   "gen:swagger",
	Short: "Generate Swagger documentation",
	Long:  `Generate Swagger documentation from annotations in code.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Kiểm tra swag CLI đã được cài đặt chưa
		_, err := exec.LookPath("swag")
		if err != nil {
			fmt.Println("Swag CLI not found. Installing...")
			installCmd := exec.Command("go", "install", "github.com/swaggo/swag/cmd/swag@latest")
			installCmd.Stdout = os.Stdout
			installCmd.Stderr = os.Stderr
			if err := installCmd.Run(); err != nil {
				return fmt.Errorf("failed to install swag: %v", err)
			}
			fmt.Println("Swag CLI installed successfully.")
		}

		// Tạo thư mục docs nếu chưa tồn tại
		docsDir := filepath.Join("docs", "swagger")
		if err := os.MkdirAll(docsDir, 0755); err != nil {
			return fmt.Errorf("failed to create docs/swagger directory: %v", err)
		}

		// Chạy swag init để tạo tài liệu Swagger
		fmt.Println("Generating Swagger documentation...")
		swagCmd := exec.Command("swag", "init",
			"--generalInfo", "main.go",
			"--output", docsDir,
			"--parseVendor",
			"--parseDependency",
		)
		swagCmd.Stdout = os.Stdout
		swagCmd.Stderr = os.Stderr
		if err := swagCmd.Run(); err != nil {
			return fmt.Errorf("failed to generate Swagger documentation: %v", err)
		}

		fmt.Println("Swagger documentation generated successfully at", docsDir)
		fmt.Println("To view the Swagger documentation, run the server and access /swagger/index.html")
		return nil
	},
}
