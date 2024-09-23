import { App, Notice, Plugin, PluginSettingTab, Setting } from "obsidian";

import JSZip from "jszip";

interface JadeHelperSettings {
	serverAddress: string;
	publisherKey: string;
}

const DEFAULT_SETTINGS: JadeHelperSettings = {
	serverAddress: "",
	publisherKey: "",
};

export default class JadeHelper extends Plugin {
	settings: JadeHelperSettings;

	async onload() {
		await this.loadSettings();

		this.addCommand({
			id: "publish-vault",
			name: "Publish vault to jade",
			callback: async () => {
				new Notice("CISSY");
				try {
					const zip = new JSZip();
					const files = this.app.vault.getFiles();

					for (const file of files) {
						const content = this.app.vault.readBinary(file);
						zip.file(file.path, content);
					}

					const zipBlob = await zip.generateAsync({ type: "blob" });

					const formData = new FormData();
					formData.append("vault", zipBlob, "vault.zip");

					const res = await fetch(
						this.settings.serverAddress + "/publish",
						{
							method: "POST",
							body: formData,
							headers: {
								PublisherKey: this.settings.publisherKey,
							},
						},
					);

					if (!res.ok) {
						throw new Error(`${res.statusText}`);
					}

					new Notice(`Published vault.`);
				} catch (error) {
					new Notice(`Error while publishing vault: ${error}`);
				}
			},
		});

		// This adds a settings tab so the user can configure various aspects of the plugin
		this.addSettingTab(new SettingsTab(this.app, this));

		// When registering intervals, this function will automatically clear the interval when the plugin is disabled.
		this.registerInterval(
			window.setInterval(() => console.log("setInterval"), 5 * 60 * 1000),
		);
	}

	onunload() { }

	async loadSettings() {
		this.settings = Object.assign(
			{},
			DEFAULT_SETTINGS,
			await this.loadData(),
		);
	}

	async saveSettings() {
		await this.saveData(this.settings);
	}
}

class SettingsTab extends PluginSettingTab {
	plugin: JadeHelper;

	constructor(app: App, plugin: JadeHelper) {
		super(app, plugin);
		this.plugin = plugin;
	}

	display(): void {
		const { containerEl } = this;

		containerEl.empty();

		new Setting(containerEl).setName("Server address").addText((text) =>
			text
				.setPlaceholder("https://my-jade-server.org")
				.setValue(this.plugin.settings.serverAddress)
				.onChange(async (value) => {
					this.plugin.settings.serverAddress = value;
					await this.plugin.saveSettings();
				}),
		);
		new Setting(containerEl).setName("Publisher key").addText((text) =>
			text
				.setValue(this.plugin.settings.publisherKey)
				.onChange(async (value) => {
					this.plugin.settings.publisherKey = value;
					await this.plugin.saveSettings();
				}),
		);
	}
}
